package service

import (
	"reflect"
	"testing"

	model "github.com/neumann-mlucas/go-url-shortener/internal/model"
	repository "github.com/neumann-mlucas/go-url-shortener/internal/repository"
)

func TestNewShortUrlService(t *testing.T) {
	type args struct {
		repository repository.ShortUrlRepository
	}
	tests := []struct {
		name string
		args args
		want *ShortUrlService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShortUrlService(tt.args.repository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortUrlService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortUrlService_CreateShortUrl(t *testing.T) {
	type fields struct {
		repository repository.ShortUrlRepository
	}
	type args struct {
		fullurl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ShortUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortUrlService{
				repository: tt.fields.repository,
			}
			got, err := s.CreateShortUrl(tt.args.fullurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortUrlService.CreateShortUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortUrlService.CreateShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortUrlService_GetShortUrl(t *testing.T) {
	type fields struct {
		repository repository.ShortUrlRepository
	}
	type args struct {
		hash string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ShortUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortUrlService{
				repository: tt.fields.repository,
			}
			got, err := s.GetShortUrl(tt.args.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortUrlService.GetShortUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortUrlService.GetShortUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortUrlService_GetShortUrls(t *testing.T) {
	type fields struct {
		repository repository.ShortUrlRepository
	}
	type args struct {
		limit int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ShortUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortUrlService{
				repository: tt.fields.repository,
			}
			got, err := s.GetShortUrls(tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShortUrlService.GetShortUrls() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ShortUrlService.GetShortUrls() = %v, want %v", got, tt.want)
			}
		})
	}
}
