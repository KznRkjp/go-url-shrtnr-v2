package db

import (
	"github.com/KznRkjp/go-url-shrtnr-v2/internal/models"
)

var urls = make(map[string]models.URL)

func SaveURL(url models.URL) {
	urls[url.Shortened] = url
}

func GetURL(shortened string) (models.URL, bool) {
	url, ok := urls[shortened]
	return url, ok
}

func PrintAllURLs() {
	for _, url := range urls {
		println(url.Original, url.Shortened)
	}
}
