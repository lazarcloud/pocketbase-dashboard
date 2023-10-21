package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/lazarcloud/pocketbase-dashboard/api/functions"
	"github.com/lazarcloud/pocketbase-dashboard/api/paths"
)

var origin string

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}

func main() {
	flag.StringVar(&origin, "origin", "http://localhost:8080/", "Define website origin")
	flag.Parse()

	fmt.Printf("Origin: %s\n", origin)

	http.HandleFunc("/", enableCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("New request: %s\n", r.URL.Path)
		pathManager := paths.NewPathManager(r)
		if pathManager.GetPartsLength() >= 3 && pathManager.GetFirstPart() == "project" {
			targetURL := "http://pocketbase-" + pathManager.Parts()[2] + ":8080"
			r.URL.Path = "/" + strings.Join(pathManager.Parts()[3:], "/")
			functions.ServeReverseProxy(targetURL, w, r)
		} else if strings.HasPrefix(r.URL.Path, "/_app/immutable/") {
			http.StripPrefix("/_app/immutable/", http.FileServer(http.Dir("/website/_app/immutable/"))).ServeHTTP(w, r)
		} else if strings.Contains(pathManager.GetFirstPart(), ".") {
			http.StripPrefix("/", http.FileServer(http.Dir("/website/"))).ServeHTTP(w, r)
		} else {
			fmt.Print("http://localhost:5173" + r.URL.Path + "\n")
			functions.ServeReverseProxy("http://localhost:5173"+r.URL.Path, w, r)
		}
	}))

	http.HandleFunc("/create", enableCORS(functions.CreateProject))
	http.HandleFunc("/projects", enableCORS(functions.GetProjects))
	http.HandleFunc("/stop", enableCORS(functions.StopProject))
	http.HandleFunc("/start", enableCORS(functions.StartProject))

	http.ListenAndServe(":80", nil)
}
