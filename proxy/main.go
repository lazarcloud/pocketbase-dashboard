package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) >= 3 && parts[1] == "dashboard" {
			slug := parts[2]
			targetURL := "http://pocketbase-" + slug + ":8080"
			r.URL.Path = "/" + strings.Join(parts[3:], "/")
			serveReverseProxy(targetURL, w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	http.ListenAndServe(":80", nil)
}

func serveReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	proxy.ServeHTTP(w, r)
}
