package server

import (
	"github.com/go-co-op/gocron"
	"time"
)

func StartHealthCheck() {
	s := gocron.NewScheduler(time.Local)
	for _, host := range serverList {
		s.Every(2).Second().Do(func(h *server) {
			h.checkHealth()
		}, host)
	}
	s.StartAsync()
}
