package main

import (
	config "github.com/neumann-mlucas/go-url-shortener/internal/config"
	handler "github.com/neumann-mlucas/go-url-shortener/internal/handler"
	repository "github.com/neumann-mlucas/go-url-shortener/internal/repository"
	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
	"net/http"
)

func main() {
	// Load config
	config.LoadConfig()

	// Initialize dependencies
	repo := repository.NewShortUrlRepository(config.AppConfig.DB)
	service := service.NewShortUrlService(repo)
	handler := handler.NewShortUrlHandler(service)

	// Create a new request multiplexer
	mux := http.NewServeMux()

	// Define routes and attach handlers
	mux.HandleFunc("GET /api/v1/{hash}", handler.GetShortUrl)
	mux.HandleFunc("POST /api/v1/{hash}", handler.CreateShortUrl)

	// Start the server
	http.ListenAndServe(":8080", mux)
}
