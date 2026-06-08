// main.go
// A simple HTTP server that returns JSON with hostname and version.
// The hostname changes per pod, demonstrating load balancing.
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	Hostname  string `json:"hostname"` // Shows which pod handled the request
	Timestamp string `json:"timestamp"`
	Version   string `json:"version"` // Read from environment variable
}

func main() {
	// APP_VERSION comes from the ConfigMap (see deployment below)
	version := os.Getenv("APP_VERSION")
	if version == "" {
		version = "1.0.0"
	}

	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		hostname, _ := os.Hostname()
		resp := Response{
			Message:   "Hello from K8s Training!",
			Hostname:  hostname,
			Timestamp: time.Now().Format(time.RFC3339),
			Version:   version,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	// Readiness: returns 200 only once the app is ready to serve traffic.
	// K8s won't send traffic to this pod until /readyz passes.
	var ready bool
	http.HandleFunc("/readyz", func(w http.ResponseWriter, r *http.Request) {
		if !ready {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprint(w, "not ready")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ready")
	})

	// Liveness: returns 200 as long as the process is alive.
	// If this fails, K8s restarts the pod.
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "alive")
	})

	// Mark the app as ready after startup is complete
	ready = true
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
