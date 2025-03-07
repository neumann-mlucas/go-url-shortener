package main

import (
	"log"
	"net/http"

	config "github.com/neumann-mlucas/go-url-shortener/internal/config"
	handler "github.com/neumann-mlucas/go-url-shortener/internal/handler"
	repository "github.com/neumann-mlucas/go-url-shortener/internal/repository"
	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
)

func main() {

	log.Println("Initalizaing App Global Config...")

	if err := config.LoadConfig(); err != nil {
		panic(err)
	}

	log.Println("Initalizaing APP services and dependencies...")

	repo := repository.NewShortUrlRepository(config.AppConfig.DB)
	service := service.NewShortUrlService(repo)

	urlHandler := handler.NewShortUrlHandler(service)
	pageHandler := handler.NewPageHandler(service)
	systemHandler := handler.NewSystemHandler()

	log.Println("Creating multiplexer and defining API routes")

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

	log.Printf("Starting server on port %s...\n", config.AppConfig.Port)

	if err := http.ListenAndServe(config.AppConfig.Port, mux); err != nil {
		panic(err)
	}
}
