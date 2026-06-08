package ctxtimeout

import (
	"context"
	"testing"
	"time"
)

func TestFetchWithTimeoutSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := FetchWithTimeout(ctx, 10*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "data" {
		t.Errorf("result = %q, want %q", result, "data")
	}
}

func TestFetchWithTimeoutExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := FetchWithTimeout(ctx, 2*time.Second)
	if err == nil {
		t.Error("expected timeout error")
	}
}
