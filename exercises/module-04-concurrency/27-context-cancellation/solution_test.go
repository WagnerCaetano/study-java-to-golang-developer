package ctxcancel

import (
	"context"
	"testing"
	"time"
)

func TestCancellableTaskFull(t *testing.T) {
	ctx := context.Background()
	result := CancellableTask(ctx, 100)
	if result != 100 {
		t.Errorf("expected 100, got %d", result)
	}
}

func TestCancellableTaskCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()
	result := CancellableTask(ctx, 100000000)
	if result >= 100000000 {
		t.Error("task should have been cancelled before completing")
	}
}
