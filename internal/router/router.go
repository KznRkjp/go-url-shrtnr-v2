package router

import (
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/app"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/logging"
	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(logging.WithLogging)
	r.Post("/", app.URLPostHandler)
	r.Get("/{id}", app.URLGetHandler) // Handle GET requests for shortened URLs
	return r
}
