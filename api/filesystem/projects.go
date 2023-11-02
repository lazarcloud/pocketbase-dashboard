package filesystem

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"

	uuid "github.com/nu7hatch/gouuid"
)

func RandomId() string {
	u, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return u.String()
}

func GetProjectsPath() string {
	return "./data/projects/"
}

type Project struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	// add other fields here
}

func GetProjects() []Project {
	var projects []Project

	// Get the path to the projects directory
	projectsPath := GetProjectsPath()

	// Read all the subdirectories within the projects directory
	files, err := os.ReadDir(projectsPath)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through all subdirectories in projects
	for _, f := range files {
		if f.IsDir() {
			// For each subdirectory, read the projectinfo.json file
			projectInfoPath := filepath.Join(projectsPath, f.Name(), "projectinfo.json")
			//check if file exists, if not pass to next iteration
			if _, err := os.Stat(projectInfoPath); os.IsNotExist(err) {
				continue
			} else {
				// log.Println("ProjectInfo file exists")
			}

			data, err := os.ReadFile(projectInfoPath)
			if err != nil {
				log.Fatal(err)
			}

			var projectInfo Project
			err = json.Unmarshal(data, &projectInfo)
			if err != nil {
				log.Fatal(err)
			}

			projects = append(projects, projectInfo)
		}
	}

	// Return the projects array
	return projects
}
func CreateProject(name string) string {
	// Generate a unique ID for the project
	id := RandomId()

	// Create a new project object
	project := Project{
		Id:   id,
		Name: name,
	}

	// Convert the project object to JSON
	projectJson, err := json.Marshal(project)
	if err != nil {
		log.Fatal(err)
	}

	// Get the path to the projects directory
	projectsPath := GetProjectsPath()

	// Create a new directory for the project
	projectPath := filepath.Join(projectsPath, id)
	err = os.Mkdir(projectPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Write the project JSON to a file in the new directory
	err = os.WriteFile(filepath.Join(projectPath, "projectinfo.json"), projectJson, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// Return the ID of the new project
	return id
}
