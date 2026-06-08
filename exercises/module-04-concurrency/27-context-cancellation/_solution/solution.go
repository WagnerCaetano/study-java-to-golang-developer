package ctxcancel

import "context"

func CancellableTask(ctx context.Context, max int) int {
	count := 0
	for i := 0; i < max; i++ {
		select {
		case <-ctx.Done():
			return count
		default:
			count++
		}
	}
	return count
}
