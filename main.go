package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	server := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(server)
	rProxy.ServeHTTP(w, r)
}

func getServer() *url.URL {
	server, _ := url.Parse("http://127.0.0.1:5000")
	return server
}
