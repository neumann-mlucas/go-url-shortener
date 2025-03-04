package service

import (
	model "github.com/neumann-mlucas/go-url-shortener/internal/model"
	repository "github.com/neumann-mlucas/go-url-shortener/internal/repository"
)

type ShortUrlService struct {
	repository repository.ShortUrlRepository
}

func NewShortUrlService(repository repository.ShortUrlRepository) *ShortUrlService {
	return &ShortUrlService{repository: repository}
}

// CreateShortUrl creates a short URL for the provided full URL and returns the created short URL.
func (s *ShortUrlService) CreateShortUrl(fullurl string) (*model.ShortUrl, error) {
	// TODO: validate / normalize URL
	// TODO: handle duplicated URL error

	id, err := s.repository.CreateShortUrl(fullurl)
	if err != nil {
		return nil, err
	}

	shorturl, err := s.repository.GetShortUrlByID(id)
	if err != nil {
		return nil, err
	}
	return shorturl, nil
}

// GetShortUrl retrieves the short URL corresponding to the provided hash.
func (s *ShortUrlService) GetShortUrl(hash string) (*model.ShortUrl, error) {
	shorturl, err := s.repository.GetShortUrlByHash(hash)
	if err != nil {
		return nil, err
	}
	return shorturl, nil
}

// GetShortUrls retrieves a list of short URLs with an optional limit on the number of results.
func (s *ShortUrlService) GetShortUrls(limit int) ([]*model.ShortUrl, error) {
	shorturls, err := s.repository.GetShortUrls(limit)
	if err != nil {
		return nil, err
	}
	return shorturls, nil
}
