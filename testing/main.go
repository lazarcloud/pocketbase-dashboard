package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	reverseProxy := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		fmt.Printf("[reverse proxy server] received request at: %s\n", time.Now())

		// Handle requests to /dashboard/<id>
		if strings.HasPrefix(req.URL.Path, "/dashboard/") {
			// id := strings.TrimPrefix(req.URL.Path, "/dashboard/")
			// Construct a new URL with the pocketbase server address
			// req.URL.Path
			//remove from path /dashboard/id
			// newPath := strings.TrimPrefix(req.URL.Path, "/dashboard/")
			// pocketbaseURL := fmt.Sprintf("http://pocketbase-lazar:8080/%s", newPath)
			// // Parse the new URL
			// proxyURL, err := url.Parse(pocketbaseURL)
			// if err != nil {
			// 	rw.WriteHeader(http.StatusInternalServerError)
			// 	_, _ = fmt.Fprint(rw, err)
			// 	return
			// }
			// // Update the request URL with the new URL
			// req.URL = proxyURL
			newPath := strings.TrimPrefix(req.URL.Path, "/dashboard/")
			originServerURL, err := url.Parse("http://pocketbase-lazar:8080")
			fmt.Println(fmt.Sprintf("http://pocketbase-lazar:8080/%s", newPath))
			if err != nil {
				log.Fatal("invalid origin server URL")
			}

			req.Host = originServerURL.Host
			req.URL.Host = originServerURL.Host
			req.URL.Scheme = originServerURL.Scheme
			req.RequestURI = ""
		} else {
			// For other requests, forward to the original server
			originServerURL, err := url.Parse("http://pocketbase:8080/")
			if err != nil {
				log.Fatal("invalid origin server URL")
			}

			req.Host = originServerURL.Host
			req.URL.Host = originServerURL.Host
			req.URL.Scheme = originServerURL.Scheme
			req.RequestURI = ""
		}

		req.Header.Set("X-Forwarded-Host", req.Host)
		req.Header.Set("X-Real-IP", req.RemoteAddr)
		req.Header.Set("X-Forwarded-For", req.RemoteAddr)
		req.Header.Set("X-Forwarded-Proto", req.URL.Scheme)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			_, _ = fmt.Fprint(rw, err)
			return
		}
		defer resp.Body.Close()

		for key, values := range resp.Header {
			for _, value := range values {
				rw.Header().Add(key, value)
			}
		}

		rw.WriteHeader(resp.StatusCode)

		_, _ = io.Copy(rw, resp.Body)
	})

	log.Fatal(http.ListenAndServe(":8080", reverseProxy))
}
