package server

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	Name   string
	URL    string
	Proxy  *httputil.ReverseProxy
	Health bool
}

var (
	serverList      []*server
	lastServedIndex = 0
)

func CreateServers(desiredCount int) {
	for i := 0; i < desiredCount; i += 1 {
		serverName := fmt.Sprintf("server-%d", i+1)
		serverUrl := fmt.Sprintf("http://localhost:500%d", i)
		serverList = append(serverList, newServer(serverName, serverUrl))
	}
}

func newServer(name string, urlStr string) *server {
	remote, err := url.Parse(urlStr)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	return &server{
		Name:   name,
		URL:    urlStr,
		Proxy:  proxy,
		Health: true,
	}
}

func (s *server) checkHealth() {
	resp, err := http.Head(s.URL)
	if err != nil {
		s.Health = false
		return
	}
	s.Health = resp.StatusCode == http.StatusOK
}

func GetHealthyServer() (*server, error) {
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
