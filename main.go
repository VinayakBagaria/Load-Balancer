package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	serverList      []*server
	serverLength    = 5
	lastServedIndex = 0
)

func main() {
	for i := 1; i <= serverLength; i += 1 {
		serverName := fmt.Sprintf("server-%d", i+1)
		serverUrl := fmt.Sprintf("http://localhost:500%d", i)
		serverList = append(serverList, newServer(serverName, serverUrl))
	}
	go startHealthCheck()
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	s, err := getHealthyServer()
	if err != nil {
		http.Error(w, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Sprintf("Processing request from %s\n", s.Name)
	s.ReverseProxy.ServeHTTP(w, r)
}

func getHealthyServer() (*server, error) {
	for i := 0; i < len(serverList); i++ {
		s := getServer()
		if s.Health {
			return s, nil
		}
	}
	return nil, fmt.Errorf("no healthy hosts")
}

func getServer() *server {
	s := serverList[lastServedIndex]
	lastServedIndex = (lastServedIndex + 1) % len(serverList)
	return s
}
