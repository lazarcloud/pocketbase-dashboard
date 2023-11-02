package docker

import (
	"os"
	"os/exec"
)

func StartCompose(path string) error {
	cmd := exec.Command("docker compose", "-f", path, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func ComposeStatus(path string) error {
	cmd := exec.Command("docker compose", "-f", path, "ps")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func KillCompose(path string) error {
	cmd := exec.Command("docker compose", "-f", path, "down")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
