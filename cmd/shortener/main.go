package main

import (
	"fmt"
	"net/http"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/config"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/flags"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/router"
	// Adjust the import path as necessary
)

// var Server = "localhost:8080"

func main() {

	flags.ParseFlags(config.Prod)
	fmt.Println("Starting URL Shortener Service...")
	r := router.NewRouter()
	fmt.Printf("Listening on http://%s", config.Prod.Server)
	err := http.ListenAndServe(config.Prod.Server, r)

	if err != nil {
		panic(err)
	}

}
