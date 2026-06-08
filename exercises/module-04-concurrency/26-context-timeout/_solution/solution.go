package ctxtimeout

import (
	"context"
	"time"
)

func FetchWithTimeout(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "data", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
