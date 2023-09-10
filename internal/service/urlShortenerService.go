package service

import (
	"crypto/md5"
	"fmt"
	"url-shortener/internal/repository"
)

type UrlShortenerService struct {
	repos repository.KeyValueRepository
}

const salt = "altcraft"

func NewUrlShortenerService(repos repository.KeyValueRepository) *UrlShortenerService {
	return &UrlShortenerService{
		repos: repos,
	}
}

func (uss *UrlShortenerService) GetUrl(hash string) (string, error) {
	val, err := uss.repos.Get(hash)
	return val, err
}

func (uss *UrlShortenerService) SaveUrl(url string) (string, error) {
	hash := encrypt(url)
	return hash, uss.repos.Set(hash, url)
}

func encrypt(url string) string {
	// hash url
	hash := md5.New()
	hash.Write([]byte(url))
	urlHash := fmt.Sprintf("%x", hash.Sum([]byte(salt)))

	// get final hash
	return urlHash[len(urlHash)-1-8 : len(urlHash)-1]
}
