package cacheaside

import (
	"context"
	"time"

	redisex "github.com/study-java-to-golang-developer/exercises/module-11-data-messaging/61-redis-get-set"
)

// GetOrLoad returns cached value or loads and stores it.
func GetOrLoad(ctx context.Context, client redisex.Client, key string, ttl time.Duration, loader func(context.Context) (string, error)) (string, error) {
	// FIX ME: Get; on ErrCacheMiss call loader then Set
	_, _ = client, loader
	return "", redisex.ErrCacheMiss
}
