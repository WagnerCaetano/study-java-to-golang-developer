package graceful

import (
	"context"
	"sync/atomic"
	"testing"
	"time"
)

func TestShutdownStopsBackgroundWork(t *testing.T) {
	srv := NewServer()

	var count atomic.Int64
	stopped := make(chan struct{})

	srv.StartBackgroundWork(func(ctx context.Context) {
		defer close(stopped)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				count.Add(1)
				time.Sleep(2 * time.Millisecond)
			}
		}
	})

	time.Sleep(15 * time.Millisecond)

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		t.Fatalf("Shutdown error: %v", err)
	}

	select {
	case <-stopped:
	case <-time.After(500 * time.Millisecond):
		t.Fatal("background work did not stop after Shutdown")
	}

	final := count.Load()
	time.Sleep(20 * time.Millisecond)
	if count.Load() != final {
		t.Fatalf("counter still increasing after shutdown: before=%d after=%d", final, count.Load())
	}
}
