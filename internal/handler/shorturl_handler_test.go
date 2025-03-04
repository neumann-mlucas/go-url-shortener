package handler

import (
	"net/http"
	"testing"

	service "github.com/neumann-mlucas/go-url-shortener/internal/service"
)

func TestShortUrlHandler_GetShortUrls(t *testing.T) {
	type fields struct {
		service *service.ShortUrlService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ShortUrlHandler{
				service: tt.fields.service,
			}
			h.GetShortUrls(tt.args.w, tt.args.r)
		})
	}
}

func TestShortUrlHandler_GetShortUrl(t *testing.T) {
	type fields struct {
		service *service.ShortUrlService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ShortUrlHandler{
				service: tt.fields.service,
			}
			h.GetShortUrl(tt.args.w, tt.args.r)
		})
	}
}

func TestShortUrlHandler_CreateShortUrl(t *testing.T) {
	type fields struct {
		service *service.ShortUrlService
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &ShortUrlHandler{
				service: tt.fields.service,
			}
			h.CreateShortUrl(tt.args.w, tt.args.r)
		})
	}
}
