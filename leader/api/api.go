package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lazarcloud/lazarbase/leader/filesystem"
)

func ServeApi() {

	r := mux.NewRouter()
	r.HandleFunc("/projects", getProjects).Methods("GET")
	r.HandleFunc("/projects/create", createProject).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	projects := filesystem.GetProjects()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}
func createProject(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var project filesystem.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the project
	id := filesystem.CreateProject(project.Name)

	// Send the ID of the new project in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}
