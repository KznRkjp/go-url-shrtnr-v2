package main

import (
	"fmt"
	"net/http"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/app" // Adjust the import path as necessary
)

var Server = "localhost:8080"

func main() {

	fmt.Println("Starting URL Shortener Service...")
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, app.URLPostHandler)
	mux.HandleFunc(`/{id}`, app.URLGetHandler)
	fmt.Printf("Listening on http://%s", Server)
	err := http.ListenAndServe(Server, mux)

	if err != nil {
		panic(err)
	}

}
