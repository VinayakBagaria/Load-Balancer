package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList = []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5001",
		"http://127.0.0.1:5002",
	}
	lastServedIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	server := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(server)
	log.Printf("Routing the request to the URL: %s", server.String())
	rProxy.ServeHTTP(w, r)
}

func getServer() *url.URL {
	server, _ := url.Parse(serverList[lastServedIndex])
	lastServedIndex = (lastServedIndex + 1) % len(serverList)
	return server
}
