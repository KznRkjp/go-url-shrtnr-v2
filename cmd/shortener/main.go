package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/config"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/flags"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/logging"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/router"
)

func main() {
	// Initialize logger
	if err := logging.InitLogger(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// Parse flags and update config
	flags.ParseFlags(config.Prod)

	fmt.Println("Starting URL Shortener Service...")
	fmt.Printf("Listening on http://%s\n", config.Prod.Server)

	// Create router
	r := router.NewRouter()

	// Start HTTP server
	if err := http.ListenAndServe(config.Prod.Server, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
