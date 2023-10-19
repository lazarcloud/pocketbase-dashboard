package main

import (
	"github.com/lazarcloud/lazarbase/leader/api"
	"github.com/lazarcloud/lazarbase/leader/filesystem"
	"github.com/lazarcloud/lazarbase/leader/implementations"
)

func main() {

	filesystem.PrepareFolderStructure()

	// docker.PrepareContainers()

	// fmt.Println("Current Container ID: ", docker.GetMyCurrentContainerId())
	filesystem.CreateProject("test4")
	implementations.PrepareSupabase("test4", "lazar376412", 3)
	// err := docker.StartCompose("data/projects/test4/docker-compose.yml")

	// fmt.Print(err.Error())

	api.ServeApi()
}
