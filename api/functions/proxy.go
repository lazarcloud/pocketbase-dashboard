package functions

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ServeReverseProxy(target string, w http.ResponseWriter, r *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	r.URL.Host = url.Host
	r.URL.Scheme = url.Scheme
	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
	r.Host = url.Host

	proxy.ServeHTTP(w, r)
}
