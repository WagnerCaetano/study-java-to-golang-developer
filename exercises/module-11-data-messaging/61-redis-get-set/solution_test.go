package redisex

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
)

type memoryClient struct {
	mu   sync.Mutex
	data map[string]entry
}

type entry struct {
	val string
	exp time.Time
}

func newMemoryClient() *memoryClient {
	return &memoryClient{data: make(map[string]entry)}
}

func (m *memoryClient) Get(ctx context.Context, key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	e, ok := m.data[key]
	if !ok || (!e.exp.IsZero() && time.Now().After(e.exp)) {
		return "", ErrCacheMiss
	}
	return e.val, nil
}

func (m *memoryClient) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	var exp time.Time
	if ttl > 0 {
		exp = time.Now().Add(ttl)
	}
	m.data[key] = entry{val: val, exp: exp}
	return nil
}

func TestCacheStoreGetSet(t *testing.T) {
	ctx := context.Background()
	store := NewCacheStore(newMemoryClient())

	if _, err := store.Get(ctx, "k"); !errors.Is(err, ErrCacheMiss) {
		t.Fatalf("Get err = %v", err)
	}
	if err := store.Set(ctx, "k", "v", time.Minute); err != nil {
		t.Fatal(err)
	}
	v, err := store.Get(ctx, "k")
	if err != nil || v != "v" {
		t.Fatalf("Get = %q, %v", v, err)
	}
}
