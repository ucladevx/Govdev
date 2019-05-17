package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(client *redis.Client) *RedisCache {
	return &RedisCache{
		client: client,
	}
}

func (rs *RedisCache) Get(key string) ([]byte, error) {
	if result, err := rs.client.Get(key).Bytes(); err != nil {
		return nil, err
	} else {
		return result, nil
	}

}

func (rs *RedisCache) Set(key string, value []byte, expiration time.Duration) error {
	if err := rs.client.Set(key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (rs *RedisCache) Del(key string) error {
	if err := rs.client.Del(key).Err(); err != nil {
		return err
	}

	return nil
}
