package mutex

import "sync"

// SafeCounter is a thread-safe counter
// FIX ME: Add mutex protection to all methods
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	// FIX: Add mutex lock/unlock
	c.count++
}

func (c *SafeCounter) Get() int {
	// FIX: Add mutex lock/unlock
	return c.count
}

func (c *SafeCounter) Reset() {
	// FIX: Add mutex lock/unlock
	c.count = 0
}
