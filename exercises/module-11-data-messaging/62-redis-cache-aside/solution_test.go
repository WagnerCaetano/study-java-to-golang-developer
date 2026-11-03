package cacheaside

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	redisex "github.com/study-java-to-golang-developer/exercises/module-11-data-messaging/61-redis-get-set"
)

type mem struct {
	mu   sync.Mutex
	data map[string]string
}

func (m *mem) Get(ctx context.Context, key string) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[key]
	if !ok {
		return "", redisex.ErrCacheMiss
	}
	return v, nil
}

func (m *mem) Set(ctx context.Context, key, val string, ttl time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
	return nil
}

func TestGetOrLoadCachesLoaderResult(t *testing.T) {
	ctx := context.Background()
	c := &mem{data: make(map[string]string)}
	var loads int
	val, err := GetOrLoad(ctx, c, "user:1", time.Minute, func(ctx context.Context) (string, error) {
		loads++
		return "alice", nil
	})
	if err != nil || val != "alice" || loads != 1 {
		t.Fatalf("first: val=%q err=%v loads=%d", val, err, loads)
	}
	val, err = GetOrLoad(ctx, c, "user:1", time.Minute, func(ctx context.Context) (string, error) {
		loads++
		return "bob", nil
	})
	if err != nil || val != "alice" || loads != 1 {
		t.Fatalf("second: val=%q err=%v loads=%d", val, err, loads)
	}
}

func TestGetOrLoadPropagatesLoaderError(t *testing.T) {
	ctx := context.Background()
	c := &mem{data: make(map[string]string)}
	errBoom := errors.New("boom")
	_, err := GetOrLoad(ctx, c, "x", time.Minute, func(ctx context.Context) (string, error) {
		return "", errBoom
	})
	if !errors.Is(err, errBoom) {
		t.Fatalf("err = %v", err)
	}
}
