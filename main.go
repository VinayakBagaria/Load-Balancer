package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverList      []*httputil.ReverseProxy
	serverLength    = 5
	lastServedIndex = 0
)

func main() {
	for i := 1; i <= serverLength; i += 1 {
		serverUrl := fmt.Sprintf("http://localhost:500%d", i)
		serverList = append(serverList, createHost(serverUrl))
	}

	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	server := getServer()
	server.ServeHTTP(w, r)
}

func getServer() *httputil.ReverseProxy {
	server := serverList[lastServedIndex]
	lastServedIndex = (lastServedIndex + 1) % len(serverList)
	return server
}

func createHost(serverUrl string) *httputil.ReverseProxy {
	server, _ := url.Parse(serverUrl)
	return httputil.NewSingleHostReverseProxy(server)
}
