package main

import (
	"fmt"
	"net/http"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/app" // Adjust the import path as necessary
)

func main() {
	server := "0.0.0.0:8080"
	fmt.Println("Starting URL Shortener Service...")
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, app.URLPostHandler)
	mux.HandleFunc(`/{id}`, app.URLGetHandler)
	fmt.Printf("Listening on http://%s", server)
	err := http.ListenAndServe(server, mux)

	if err != nil {
		panic(err)
	}

}
