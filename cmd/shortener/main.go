package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/config"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/flags"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/logging"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/router"
	// Adjust the import path as necessary
)

// var Server = "localhost:8080"

func main() {
	if err := logging.InitLogger(); err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	flags.ParseFlags(config.Prod)
	fmt.Println("Starting URL Shortener Service...")
	r := router.NewRouter()
	fmt.Printf("Listening on http://%s\n", config.Prod.Server)
	err := http.ListenAndServe(config.Prod.Server, r)

	if err != nil {
		panic(err)
	}

}
