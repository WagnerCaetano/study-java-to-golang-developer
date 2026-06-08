package ctxtimeout

import (
	"context"
	"fmt"
	"time"
)

// FetchWithTimeout simulates a fetch that respects context timeout
// FIX ME: Return error if context times out
func FetchWithTimeout(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "data", nil
	// FIX: Add case for ctx.Done()
	}
}
