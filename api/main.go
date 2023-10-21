package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lazarcloud/pocketbase-dashboard/api/functions"
	"github.com/lazarcloud/pocketbase-dashboard/api/paths"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("New request: %s\n", r.URL.Path)
		pathManager := paths.NewPathManager(r)
		if pathManager.GetPartsLength() >= 3 && pathManager.GetFirstPart() == "dashboard" {
			targetURL := "http://pocketbase-" + pathManager.Parts()[2] + ":8080"
			r.URL.Path = "/" + strings.Join(pathManager.Parts()[3:], "/")
			functions.ServeReverseProxy(targetURL, w, r)
		} else if strings.HasPrefix(r.URL.Path, "/_app/immutable/") {
			http.StripPrefix("/_app/immutable/", http.FileServer(http.Dir("/_app/immutable/"))).ServeHTTP(w, r)
		} else {
			fmt.Print("http://localhost:5173" + r.URL.Path)
			functions.ServeReverseProxy("http://localhost:5173"+r.URL.Path, w, r)
		}
	})

	http.HandleFunc("/create", functions.CreateProject)

	http.HandleFunc("/containers", functions.GetProjects)

	http.HandleFunc("/stop", functions.StopProject)
	http.HandleFunc("/start", functions.StartProject)

	http.ListenAndServe(":80", nil)
}
