package main

import (
	"log"
	"os"

	"github.com/msalahm24/simple-http-api/server"
)

func main() {

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	srv := server.NewServer(port)

	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
