package main

import (
	"fmt"
	"github.com/VinayakBagaria/load-balancer/server"
	"log"
	"net/http"
)

func main() {
	server.CreateServers(5)
	server.StartHealthCheck()
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(w http.ResponseWriter, r *http.Request) {
	s, err := server.GetHealthyServer()
	if err != nil {
		http.Error(w, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	fmt.Sprintf("Processing request from %s\n", s.Name)
	s.ReverseProxy.ServeHTTP(w, r)
}
