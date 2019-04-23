package redis

import (
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v6"
)

func NewConnection(host, port, password string, db int) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       db,
	})

	if _, err := client.Ping().Result(); err != nil {
		fmt.Println("Error connecting to cache", err)
		os.Exit(1)
	}

	return client
}

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(client *redis.Client) *RedisStore {
	return &RedisStore{
		client: client,
	}
}

func (rs *RedisStore) Get(key string) ([]byte, error) {
	if result, err := rs.client.Get(key).Bytes(); err != nil {
		return nil, err
	} else {
		return result, nil
	}

}

func (rs *RedisStore) Set(key string, value []byte, expiration time.Duration) error {
	if err := rs.client.Set(key, value, expiration).Err(); err != nil {
		return err
	}
	return nil
}

func (rs *RedisStore) Del(key string) error {
	if err := rs.client.Del(key).Err(); err != nil {
		return err
	}

	return nil
}
