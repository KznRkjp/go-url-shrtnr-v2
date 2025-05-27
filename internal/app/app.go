package app

import (
	"fmt"
	"net/http"
)

func URLPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Handle URL shortening logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL shortened successfully"))
}
func URLGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Extract the ID from the URL path
	id := r.URL.Path[len("/"):] // Assuming the path is like /{id}
	fmt.Println(r.URL.Path)
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	// Here you would typically look up the ID in your database or storage
	// Handle URL retrieval logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(id))
}
