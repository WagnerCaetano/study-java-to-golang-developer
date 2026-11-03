package timeoutpat

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRunWithTimeoutFastWork(t *testing.T) {
	err := RunWithTimeout(context.Background(), 200*time.Millisecond, func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunWithTimeoutSlowWork(t *testing.T) {
	err := RunWithTimeout(context.Background(), 20*time.Millisecond, func(ctx context.Context) error {
		select {
		case <-time.After(500 * time.Millisecond):
			return nil
		case <-ctx.Done():
			return ctx.Err()
		}
	})
	if err == nil {
		t.Fatal("expected timeout error, got nil")
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context.DeadlineExceeded", err)
	}
}
