package service

import (
	"crypto/md5"
	"fmt"
)

const salt = "altcraft"

type Hasher interface {
	Encrypt(text string) string
}

type Md5Hasher struct {
}

func NewMd5Hasher() *Md5Hasher {
	return &Md5Hasher{}
}

func (h *Md5Hasher) Encrypt(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	urlHash := fmt.Sprintf("%x", hash.Sum([]byte(salt)))
	return urlHash[len(urlHash)-1-8 : len(urlHash)-1]
}
