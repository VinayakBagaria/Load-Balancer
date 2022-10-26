package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func newServer(name string, urlStr string) *server {
	u, _ := url.Parse(urlStr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		Name:         name,
		URL:          urlStr,
		ReverseProxy: rp,
		Health:       true,
	}
}

func (s *server) checkHealth() {
	resp, err := http.Head(s.URL)
	s.Health = err != nil && resp.StatusCode == http.StatusOK
}
