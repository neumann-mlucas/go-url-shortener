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

func (s *ShortUrlService) CreateShortUrl(fullurl string) error {
	// TODO: should handle duplicated url error?
	return s.repository.CreateShortUrl(fullurl)
}

func (s *ShortUrlService) GetShortUrl(hash string) (*model.ShortUrl, error) {
	shorturl, err := s.repository.GetShortUrlByHash(hash)
	if err != nil {
		return nil, err
	}
	return shorturl, nil
}

func (s *ShortUrlService) GetShortUrls(limit int) ([]*model.ShortUrl, error) {
	shorturls, err := s.repository.GetShortUrls(limit)
	if err != nil {
		return nil, err
	}
	return shorturls, nil
}
