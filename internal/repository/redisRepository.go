package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisRepository struct {
	client  *redis.Client
	context *context.Context
}

func NewRedisRepository(client *redis.Client, context *context.Context) *RedisRepository {
	return &RedisRepository{
		client:  client,
		context: context,
	}
}

func (r *RedisRepository) Get(key string) (string, error) {
	logrus.Debugf("RedisRepository Get k=(%s)", key)
	if val, err := r.client.Get(*r.context, key).Result(); err != nil {
		return "", err
	} else {
		return val, nil
	}
}

func (r *RedisRepository) Set(key, value string) error {
	logrus.Debugf("RedisRepository Set k=(%s), v=(%s)", key, value)
	if err := r.client.Set(*r.context, key, value, 0).Err(); err != nil {
		return err
	}
	return nil
}
