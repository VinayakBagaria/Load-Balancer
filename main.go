package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList      []string
	serverLength    = 5
	lastServedIndex = 0
)

func main() {
	for i := 1; i <= serverLength; i += 1 {
		serverUrl := fmt.Sprintf("http://localhost:500%d", i)
		serverList = append(serverList, serverUrl)
	}

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
