package service

import (
	"internal/repository"
	"internal/utils"
)

type ShortUrlService struct {
	repository repository.ShortUrlRepository
}

func NewShortUrlService(repository repository.ShortUrlRepository) *ShortUrlService {
	return &ShortUrlService{repository: repository}
}

func (s *ShortUrlService) CreateShortUrl(fullurl string) {
	err := s.repository.CreateShortUrl(fullurl)
	if err != nil {
		return nil, err
	}
}

func (s *ShortUrlService) GetShortUrl(hash string) {
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
