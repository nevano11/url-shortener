package service

import (
	"url-shortener/internal/repository"
)

type UrlShortenerService struct {
	repos  repository.KeyValueRepository
	hasher Hasher
}

func NewUrlShortenerService(repos repository.KeyValueRepository, hasher Hasher) *UrlShortenerService {
	return &UrlShortenerService{
		repos:  repos,
		hasher: hasher,
	}
}

func (uss *UrlShortenerService) GetUrl(hash string) (string, error) {
	val, err := uss.repos.Get(hash)
	return val, err
}

func (uss *UrlShortenerService) SaveUrl(url string) (string, error) {
	hash := uss.hasher.Encrypt(url)
	return hash, uss.repos.Set(hash, url)
}
