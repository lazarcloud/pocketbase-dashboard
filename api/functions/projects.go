package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/lazarcloud/pocketbase-dashboard/api/globals"
)

type ContainerDetails struct {
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	// Get list of all containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter containers whose names start with "pocketbase-"
	var pocketbaseContainers []map[string]interface{}
	for _, container := range containers {
		if strings.HasPrefix(container.Names[0], "/pocketbase-") && container.NetworkSettings.Networks["lazar-static"] != nil {
			containerInfo := map[string]interface{}{
				"Name":   container.Names[0][1:], // Remove leading '/'
				"Status": container.State,
			}

			// Check if details file exists, if not, create one with default values
			detailsFilePath := filepath.Join(globals.DataFolder, container.Names[0][1:]+"_details.json")
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
			detailsBytes, err := os.ReadFile(detailsFilePath)
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
}
func CreateProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("slug")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "monsieurlazar/pocketbase",
		},
		nil, nil, nil, "pocketbase-"+id,
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
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	//return resp.ID to client answer to api
	fmt.Println(resp.ID)

	w.Write([]byte(resp.ID))
}
func StopProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("slug")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// Stop the container with the given ID
	err = cli.ContainerStop(context.Background(), "pocketbase-"+id, container.StopOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Container stopped successfully"))
}

func StartProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("slug")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	// Start the container with the given ID
	err = cli.ContainerStart(context.Background(), "pocketbase-"+id, types.ContainerStartOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Container started successfully"))
}
