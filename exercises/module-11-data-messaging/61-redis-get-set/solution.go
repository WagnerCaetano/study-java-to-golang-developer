package redisex

import (
	"context"
	"errors"
	"time"
)

// ErrCacheMiss is returned when a key is not present or expired.
var ErrCacheMiss = errors.New("cache miss")

// Client is the subset of go-redis used by CacheStore.
type Client interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, val string, ttl time.Duration) error
}

// CacheStore provides cache get/set on top of Client.
type CacheStore struct {
	client Client
}

// NewCacheStore wraps a Client.
func NewCacheStore(c Client) *CacheStore {
	return &CacheStore{client: c}
}

// Get returns a cached value.
func (s *CacheStore) Get(ctx context.Context, key string) (string, error) {
	// FIX ME: delegate to client; map empty/missing to ErrCacheMiss
	_, _ = ctx, key
	return "", ErrCacheMiss
}

// Set stores a value with TTL.
func (s *CacheStore) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	// FIX ME: delegate to client
	return errors.New("not implemented")
}
