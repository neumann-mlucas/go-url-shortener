package main

import (
	"internal/config"
	"internal/handler"
	"internal/middleware"
	"internal/repository"
	"internal/service"
	"net/http"
)

func main() {
	// Load config
	config.LoadConfig()

	// Initialize dependencies
	repo := repository.NewUrlRepository(config.AppConfig.DB)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(service)

	// Create a new request multiplexer
	mux := http.NewServeMux()

	// Define routes and attach handlers
	mux.HandleFunc("GET /api/v1/{hash}", handler.GetShortUrl)
	mux.HandleFunc("POST /api/v1/{hash}", handler.CreateShortUrl)

	// Start the server
	http.ListenAndServe(":8080", mux)
}
