package handler

import (
	"encoding/json"
	"net/http"

	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
)

type ShortUrlHandler struct {
	service *service.ShortUrlService
}

func NewShortUrlHandler(service *service.ShortUrlService) *ShortUrlHandler {
	return &ShortUrlHandler{service: service}
}

func (h *ShortUrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get Path Value
	fullurl := r.PathValue("url")

	// TODO: Try Retrieve Short Url

	// Create Short Url
	shorturl, err := h.service.CreateShortUrl(fullurl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write Short Url to Response Body
	json.NewEncoder(w).Encode(shorturl)

	// Write HTTP Status
	w.WriteHeader(http.StatusCreated)
}

func (h *ShortUrlHandler) GetShortUrl(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get Path Value
	hash := r.PathValue("hash")

	// Retrieve Short Url
	shorturl, err := h.service.GetShortUrl(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write Short Url to Response Body
	json.NewEncoder(w).Encode(shorturl)

	// Write HTTP Status
	w.WriteHeader(http.StatusAccepted)
}
