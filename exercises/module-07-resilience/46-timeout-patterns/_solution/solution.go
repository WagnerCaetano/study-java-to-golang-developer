package timeoutpat

import (
	"context"
	"time"
)

func RunWithTimeout(ctx context.Context, d time.Duration, work func(context.Context) error) error {
	ctx, cancel := context.WithTimeout(ctx, d)
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- work(ctx)
	}()

	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
