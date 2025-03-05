package main

import (
	"fmt"
	"net/http"

	config "github.com/neumann-mlucas/go-url-shortener/internal/config"
	handler "github.com/neumann-mlucas/go-url-shortener/internal/handler"
	repository "github.com/neumann-mlucas/go-url-shortener/internal/repository"
	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
)

func main() {
	// Load config
	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	// Initialize dependencies
	repo := repository.NewShortUrlRepository(config.AppConfig.DB)
	service := service.NewShortUrlService(repo)

	urlHandler := handler.NewShortUrlHandler(service)
	pageHandler := handler.NewPageHandler(service)
	systemHandler := handler.NewSystemHandler()

	// Create a new request multiplexer
	mux := http.NewServeMux()

	// Define API routes
	mux.HandleFunc("GET  /api/url/{hash}", urlHandler.GetShortUrl)
	mux.HandleFunc("GET  /api/url", urlHandler.GetShortUrls)
	mux.HandleFunc("POST /api/url", urlHandler.CreateShortUrl)

	// Define USER routes
	mux.HandleFunc("GET /", pageHandler.ServeLandingPage)
	mux.HandleFunc("GET /{hash}", pageHandler.RedirectShortUrl)

	// Define SYSTEM / ADMIN routes
	mux.HandleFunc("GET /health", systemHandler.HealthCheck)
	mux.HandleFunc("GET /doc", systemHandler.RedirectDocs)

	// Start the server
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
