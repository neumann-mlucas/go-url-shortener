package service

import (
	model "github.com/neumann-mlucas/go-url-shortener/internal/model"
	repository "github.com/neumann-mlucas/go-url-shortener/internal/repository"
	utils "github.com/neumann-mlucas/go-url-shortener/internal/utils"
)

type ShortUrlService struct {
	repository *repository.ShortUrlRepository
}

func NewShortUrlService(repository *repository.ShortUrlRepository) *ShortUrlService {
	return &ShortUrlService{repository: repository}
}

func (s *ShortUrlService) CreateShortUrl(fullurl string) error {
	err := s.repository.CreateShortUrl(fullurl)
	if err != nil {
		return err
	}
}

func (s *ShortUrlService) GetShortUrl(hash string) (*model.ShortUrl, error) {
	id, err := utils.ToID(hash)
	if err != nil {
		return nil, err
	}

	shorturl, err := s.repository.GetShortUrlByID(id)
	if err != nil {
		return nil, err
	}

	return shorturl
}
