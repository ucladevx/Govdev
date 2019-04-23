package services

import (
	"time"
)

type CacheStore interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte, expiration time.Duration) error
	Del(key string) error
}

type CacheService struct {
	store CacheStore
}

func NewCacheService(store CacheStore) *CacheService {
	return &CacheService{
		store: store,
	}

}

func (cs *CacheService) Get(key string) ([]byte, error) {
	return cs.store.Get(key)
}

func (cs *CacheService) Set(key string, value []byte, expiration time.Duration) error {
	return cs.store.Set(key, value, expiration)

}

func (cs *CacheService) Del(key string) error {
	return cs.store.Del(key)
}
