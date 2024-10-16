package main

import (
	"event-router/src/app/handlers/event"
	"event-router/src/app/internal/config"
	"net/http"
	"os"

	"fmt"

	"github.com/sirupsen/logrus"
	// "github.com/segmentio/kafka-go"
)

func main() {
	// Configure logging
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Error("Failed to read config.", err)
		return
	}

	http.HandleFunc("/event", event.HandleEvent)
	http.HandleFunc("/event/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			event.HandleEvent(w, r)
			return
		}
		w.WriteHeader(http.StatusMethodNotAllowed)
	})

	addr := fmt.Sprintf("%s:%d", cfg.Service.Host, cfg.Service.Port)
	logrus.Infof("Service is running on %s", addr)
	logrus.Fatal(http.ListenAndServe(addr, nil))
}
