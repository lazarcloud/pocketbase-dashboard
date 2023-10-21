package functions

import (
	"context"
	"encoding/json"
	"fmt"
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

func respondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func GetProjects(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error creating Docker client",
			"error":   err.Error(),
		})
		return
	}
	// Get list of all containers
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error getting container list",
			"error":   err.Error(),
		})
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
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error converting container list to JSON",
			"error":   err.Error(),
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}
func CreateProject(w http.ResponseWriter, r *http.Request) {

	id := r.URL.Query().Get("slug")
	if id == "" {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Provide a slug wirh ?slug=<container-name>",
			"error":   "No slug provided",
		})
		return
	}

	detailsFilePath := filepath.Join(globals.DataFolder, "pocketbase-"+id+"_details.json")
	_, err := os.Stat(detailsFilePath)
	if os.IsExist(err) {
		os.Remove(detailsFilePath)
	}

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error creating Docker client",
			"error":   err.Error(),
		})
		return
	}

	volumeMapping := fmt.Sprintf("/home/pocketbase/projects/%s:/pb/pb_data", id)
	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: "monsieurlazar/pocketbase",
		},
		&container.HostConfig{
			Binds: []string{volumeMapping},
		}, nil, nil, "pocketbase-"+id,
	)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error creating container",
			"error":   err.Error(),
		})
		return
	}

	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error starting container",
			"error":   err.Error(),
		})
		return
	}

	err = cli.NetworkConnect(context.Background(), "lazar-static", resp.ID, nil)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error connecting container to network",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println(resp.ID)
	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Container created successfully",
		"id":      resp.ID,
	})
}

func StopProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("slug")
	if id == "" {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Provide a slug wirh ?slug=<container-name>",
			"error":   "No slug provided",
		})
		return
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error creating Docker client",
			"error":   err.Error(),
		})
		return
	}

	// Stop the container with the given ID
	err = cli.ContainerStop(context.Background(), "pocketbase-"+id, container.StopOptions{})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error stopping container",
			"error":   err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Container stopped successfully",
		"id":      "pocketbase-" + id,
	})

}

func StartProject(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("slug")
	if id == "" {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Provide a slug wirh ?slug=<container-name>",
			"error":   "No slug provided",
		})
		return
	}
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error creating Docker client",
			"error":   err.Error(),
		})
		return
	}

	// Start the container with the given ID
	err = cli.ContainerStart(context.Background(), "pocketbase-"+id, types.ContainerStartOptions{})
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{
			"message": "Error starting container",
			"error":   err.Error(),
		})
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Container started successfully",
		"id":      "pocketbase-" + id,
	})
}
