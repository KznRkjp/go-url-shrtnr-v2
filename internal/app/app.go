package app

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/config"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/db"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/models"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/urlgen"
)

func URLPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	if len(body) == 0 {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}
	// Here you would typically save the URL to your database or storage
	shortURL := saveURL(string(body))
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("%s/%s", config.Prod.ServerResponse, shortURL)))

}
func URLGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Extract the ID from the URL path
	id := r.URL.Path[len("/"):] // Assuming the path is like /{id}
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	// Here you would typically look up the ID in your database or storage
	// Handle URL retrieval logic here
	urlData, ok := db.GetURL(id)
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}
	redirect(urlData, w, r)
}

func saveURL(url string) string {
	shortened := urlgen.GenerateShortKey()
	urlData := models.URL{
		Original:  url,
		Shortened: shortened, // Example shortened URL
		CreatedAt: time.Now(),
	}
	db.SaveURL(urlData) // Example timestamp
	return shortened
}

func redirect(urlData models.URL, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, urlData.Original, http.StatusTemporaryRedirect)
}
