package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

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
