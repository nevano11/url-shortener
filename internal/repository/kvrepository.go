package repository

import (
	"errors"
	"github.com/sirupsen/logrus"
)

type KeyValueRepository interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type FakeKVRepository struct {
	storage map[string]string
}

func NewFakeKVRepository() *FakeKVRepository {
	return &FakeKVRepository{
		storage: make(map[string]string),
	}
}

func (r *FakeKVRepository) Get(key string) (string, error) {
	logrus.Debugf("FakeKVRepository Get k=(%s)", key)
	val, isExists := r.storage[key]
	if !isExists {
		return "", errors.New("item not founded")
	}
	return val, nil
}

func (r *FakeKVRepository) Set(key, value string) error {
	logrus.Debugf("FakeKVRepository Set k=(%s), v=(%s)", key, value)
	r.storage[key] = value
	return nil
}
