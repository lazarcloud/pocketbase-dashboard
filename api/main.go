package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

const dataFolder = "./data"

type ContainerDetails struct {
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func main() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		fmt.Printf("New request: %s\n", r.URL.Path)

		if len(parts) >= 3 && parts[1] == "dashboard" {
			slug := parts[2]
			targetURL := "http://pocketbase-" + slug + ":8080"
			r.URL.Path = "/" + strings.Join(parts[3:], "/")
			serveReverseProxy(targetURL, w, r)
		} else if strings.HasPrefix(r.URL.Path, "/_app/immutable/") {
			// Serve static files from http://localhost:5173/_app/immutable/
			http.StripPrefix("/_app/immutable/", http.FileServer(http.Dir("/_app/immutable/"))).ServeHTTP(w, r)
		} else if r.URL.Path == "/docker" {
			// Create a container
			resp, err := cli.ContainerCreate(
				context.Background(),
				&container.Config{
					Image: "monsieurlazar/pocketbase",
				},
				nil, nil, nil, "pocketbase-"+r.URL.Query().Get("slug"),
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			//start container
			err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}

			//join network
			err = cli.NetworkConnect(context.Background(), "lazar-static", resp.ID, nil)

			//return resp.ID to client answer to api
			fmt.Println(resp.ID)

			w.Write([]byte(resp.ID))

		} else {
			fmt.Print("http://localhost:5173" + r.URL.Path)
			serveReverseProxy("http://localhost:5173"+r.URL.Path, w, r)
		}
	})

	http.HandleFunc("/containers", func(w http.ResponseWriter, r *http.Request) {
		// Get list of all containers
		containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Filter containers whose names start with "pocketbase-"
		var pocketbaseContainers []map[string]interface{}
		for _, container := range containers {
			if strings.HasPrefix(container.Names[0], "/pocketbase-") {
				containerInfo := map[string]interface{}{
					"Name":   container.Names[0][1:], // Remove leading '/'
					"Status": container.State,
				}

				// Check if details file exists, if not, create one with default values
				detailsFilePath := filepath.Join(dataFolder, container.Names[0][1:]+"_details.json")
				_, err := os.Stat(detailsFilePath)
				if os.IsNotExist(err) {
					details := ContainerDetails{
						CreatedAt:   time.Now(),
						Name:        container.Names[0][1:],
						Description: "Default description",
					}
					jsonBytes, _ := json.Marshal(details)
					os.WriteFile(detailsFilePath, jsonBytes, 0644)
				}

				// Read details from the file and append to the response
				detailsBytes, err := ioutil.ReadFile(detailsFilePath)
				if err == nil {
					var details ContainerDetails
					json.Unmarshal(detailsBytes, &details)
					containerInfo["Details"] = details
				}

				pocketbaseContainers = append(pocketbaseContainers, containerInfo)
			}
		}

		// Convert the container list to JSON and send it as the response
		jsonBytes, err := json.Marshal(pocketbaseContainers)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBytes)
	})

	http.ListenAndServe(":80", nil)
}

func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	proxy.ServeHTTP(w, r)
}
