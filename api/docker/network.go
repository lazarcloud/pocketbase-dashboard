package docker

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/lazarcloud/lazarbase/leader/globals"
)

func CheckNetwork() {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}

	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	exists := false

	for _, network := range networks {
		if network.Name == globals.NetworkName {
			exists = true
		}
	}

	if !exists {
		fmt.Printf("Created new network with name: %s\n", globals.NetworkName)
		CreateNetwork()
	} else {
		fmt.Printf("Network with name: %s already exists\n", globals.NetworkName)
	}
}

func EnsureContainerInNetwork(containerID string, networkID string) error {

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		log.Fatal(err)
	}
	// Get the container's information
	containerJSON, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return err
	}

	// Check if the container is already in the network
	for _, network := range containerJSON.NetworkSettings.Networks {
		if network.NetworkID == networkID {
			fmt.Printf("Container is already in network: %s\n", networkID)
			return nil
		}
	}

	// If not, connect the container to the network
	err = cli.NetworkConnect(ctx, networkID, containerID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Connected container: %s to network: %s\n", containerID, networkID)
	return nil
}
