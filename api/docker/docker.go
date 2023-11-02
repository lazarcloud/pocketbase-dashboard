package docker

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/lazarcloud/lazarbase/leader/globals"
)

func PrepareContainers() {

	CheckDashboard()

	CheckNetwork()

	containerID, err := GetMyCurrentContainerId()
	if err != nil {
		log.Fatal(err)
	}

	err = EnsureContainerInNetwork(containerID, globals.NetworkName)
	if err != nil {
		log.Fatal(err)
	}

}

func PullImage(image string) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	out, err := cli.ImagePull(context.Background(), image, types.ImagePullOptions{})
	if err != nil {
		log.Fatal(err)
	}

	defer out.Close()
}

func CheckDashboard() {

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	exists := false

	for _, container := range containers {
		fmt.Println(container.Names)
		if container.Names[0] == "/lazarbase-dashboard" {
			exists = true
		}
	}

	if !exists {
		fmt.Println("Creating new container")
		CreateDashboard()
	} else {
		fmt.Println("Container already exists")
	}

}

func CreateNetwork() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := cli.NetworkCreate(context.Background(), "lazarbase-network", types.NetworkCreate{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.ID)
}
func CreateDashboard() {
	// create a new container with name lazarbase-dashboard
	// pullImage("lazarbase-dashboard")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Image: globals.DashboardImage,
			ExposedPorts: nat.PortSet{
				"80/tcp": struct{}{},
			},
		},
		&container.HostConfig{
			PortBindings: nat.PortMap{
				"80/tcp": []nat.PortBinding{
					{
						HostIP:   "0.0.0.0",
						HostPort: "8081",
					},
				},
			},
			NetworkMode: container.NetworkMode("lazarbase-network"),
		}, nil, nil, "lazarbase-dashboard",
	)
	if err != nil {
		log.Fatal(err)
	}

	err = cli.NetworkConnect(context.Background(), "lazarbase-network", resp.ID, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Start the container
	err = cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetContainers(w http.ResponseWriter, r *http.Request) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, container := range containers {
		fmt.Fprintln(w, container.ID)
	}
}
func GetMyCurrentContainerId() (string, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return "", err
	}

	container, err := cli.ContainerInspect(context.Background(), "/lazarbase-leader")
	if err != nil {
		return "", err
	}

	return container.ID, nil
}
