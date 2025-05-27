package app

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/KznRkjp/go-url-shrtnr-v2/internal/db"
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/models"
)

func TestURLPostHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name:           "Valid URL",
			requestBody:    "https://example.com",
			expectedStatus: http.StatusCreated, // Change to expected status
		},
		{
			name:           "Empty Body",
			requestBody:    "",
			expectedStatus: http.StatusBadRequest, // Change to expected status
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.requestBody))
			w := httptest.NewRecorder()

			URLPostHandler(w, req)

			resp := w.Result()
			defer resp.Body.Close()

			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
			}
			// Optionally, check response body here
		})
	}
}

func TestURLGetHandler(t *testing.T) {
	// Prepare test data
	originalURL := "https://example.com"
	shortKey := "testkey"
	urlData := models.URL{
		Original:  originalURL,
		Shortened: shortKey,
		CreatedAt: time.Now(),
	}
	db.SaveURL(urlData)

	t.Run("Valid short key", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/"+shortKey, nil)
		w := httptest.NewRecorder()

		URLGetHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusTemporaryRedirect {
			t.Errorf("expected status %d, got %d", http.StatusTemporaryRedirect, resp.StatusCode)
		}
		location := resp.Header.Get("Location")
		if location != originalURL {
			t.Errorf("expected redirect to %s, got %s", originalURL, location)
		}
	})

	t.Run("Missing ID", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()

		URLGetHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/notfound", nil)
		w := httptest.NewRecorder()

		URLGetHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("expected status %d, got %d", http.StatusNotFound, resp.StatusCode)
		}
	})

	t.Run("Wrong method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/"+shortKey, nil)
		w := httptest.NewRecorder()

		URLGetHandler(w, req)

		resp := w.Result()
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("expected status %d, got %d", http.StatusMethodNotAllowed, resp.StatusCode)
		}
	})
}
