package timeoutpat

import (
	"context"
	"time"
)

// RunWithTimeout runs work with a timeout of d.
// FIX ME: Apply timeout to context and return ctx.Err() when the deadline is exceeded.
func RunWithTimeout(ctx context.Context, d time.Duration, work func(context.Context) error) error {
	_ = d
	return work(ctx)
}
