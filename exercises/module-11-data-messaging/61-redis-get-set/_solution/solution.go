package redisex

import (
	"context"
	"errors"
	"time"
)

var ErrCacheMiss = errors.New("cache miss")

type Client interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

type CacheStore struct {
	client Client
}

func NewCacheStore(c Client) *CacheStore {
	return &CacheStore{client: c}
}

func (s *CacheStore) Get(ctx context.Context, key string) (string, error) {
	return s.client.Get(ctx, key)
}

func (s *CacheStore) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	return s.client.Set(ctx, key, val, ttl)
}
