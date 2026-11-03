package cacheaside

import (
	"context"
	"errors"
	"time"

	redisex "github.com/study-java-to-golang-developer/exercises/module-11-data-messaging/61-redis-get-set"
)

func GetOrLoad(ctx context.Context, client redisex.Client, key string, ttl time.Duration, loader func(context.Context) (string, error)) (string, error) {
	val, err := client.Get(ctx, key)
	if err == nil {
		return val, nil
	}
	if !errors.Is(err, redisex.ErrCacheMiss) {
		return "", err
	}
	val, err = loader(ctx)
	if err != nil {
		return "", err
	}
	if err := client.Set(ctx, key, val, ttl); err != nil {
		return "", err
	}
	return val, nil
}
