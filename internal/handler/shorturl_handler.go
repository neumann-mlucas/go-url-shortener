package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
)

type CreateShortUrlRequest struct {
	FullUrl string `json:"full_url"`
}

type ShortUrlResponse struct {
	ShortUrl string `json:"short_url"`
}

type ShortUrlHandler struct {
	service *service.ShortUrlService
}

func NewShortUrlHandler(service *service.ShortUrlService) *ShortUrlHandler {
	return &ShortUrlHandler{service: service}
}

// CreateShortUrl handles the creation of a short URL and responds with a status.
func (h *ShortUrlHandler) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get JSON payload
	var payload CreateShortUrlRequest

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Create Short Url
	shorturl, err := h.service.CreateShortUrl(payload.FullUrl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write HTTP Status
	w.WriteHeader(http.StatusCreated)

	// Encode the response as JSON and send it back to the client
	response := ShortUrlResponse{
		ShortUrl: fmt.Sprintf("http://%s/%s", r.Host, shorturl.Hash),
	}
	json.NewEncoder(w).Encode(response)
}

// GetShortUrl retrieves a short URL based on the hash and responds with the corresponding full URL.
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

	// Write HTTP Status
	w.WriteHeader(http.StatusAccepted)

	// Encode the response as JSON and send it back to the client
	response := ShortUrlResponse{
		ShortUrl: fmt.Sprintf("http://%s/%s", r.Host, shorturl.Hash),
	}
	json.NewEncoder(w).Encode(response)

}

// GetShortUrls retrieves a list of short URLs with an optional limit and responds with the results.
func (h *ShortUrlHandler) GetShortUrls(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Get the query parameters from the URL
	params := r.URL.Query()
	limitStr := params.Get("limit")

	// Default value for limit
	limit := 1000

	// Convert limit from string to int using strconv.Atoi
	parsedLimit, err := strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid 'limit' parameter", http.StatusBadRequest)
		return
	}
	limit = parsedLimit

	// Retrieve Short Urls
	shorturls, err := h.service.GetShortUrls(limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write Short Url to Response Body
	json.NewEncoder(w).Encode(shorturls)
}
