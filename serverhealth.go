package main

import (
	"log"
	"net/http"
	"time"
)

// Configuration constants
const (
	serviceURL    = "http://example.com/health"
	checkInterval = 30 * time.Second
)

// checkServiceHealth sends a GET request to the service URL and logs the status.
func checkServiceHealth() {
	resp, err := http.Get(serviceURL)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Printf("Service is healthy: %s\n", resp.Status)
	} else {
		log.Printf("Service is not healthy: %s\n", resp.Status)
	}
}

func main() {
	log.Println("Starting service health checker...")

	for {
		checkServiceHealth()
		time.Sleep(checkInterval)
	}
}
