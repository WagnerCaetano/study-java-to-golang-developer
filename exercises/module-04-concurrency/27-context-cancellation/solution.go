package ctxcancel

import "context"

// CancellableTask runs until cancelled or count reaches max
// FIX ME: Check ctx.Done() in the loop
func CancellableTask(ctx context.Context, max int) int {
	count := 0
	for i := 0; i < max; i++ {
		count++
		// FIX: Check for cancellation
	}
	return count
}
