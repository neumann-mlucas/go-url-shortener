package handler

import (
	"net/http"

	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
)

type PageHandler struct {
	service *service.ShortUrlService
}

func NewPageHandler(service *service.ShortUrlService) *PageHandler {
	return &PageHandler{service: service}
}

func (h *PageHandler) ServeLandingPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func (h *PageHandler) RedirectShortUrl(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	shorturl, err := h.service.GetShortUrl(hash)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, shorturl.Url, http.StatusFound)
}
