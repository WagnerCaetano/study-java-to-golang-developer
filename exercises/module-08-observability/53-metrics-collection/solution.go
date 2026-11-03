package metrics

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

// Counter counts events with optional label pairs.
type Counter interface {
	Inc(labels ...string)
}

// MapCounter is an in-memory labelled counter for tests and local dev.
type MapCounter struct {
	mu     sync.Mutex
	values map[string]int64
}

func NewMapCounter() *MapCounter {
	return &MapCounter{values: make(map[string]int64)}
}

func (c *MapCounter) key(labels ...string) string {
	// labels stored as joined pairs
	s := ""
	for i := 0; i+1 < len(labels); i += 2 {
		s += labels[i] + "=" + labels[i+1] + "|"
	}
	return s
}

func (c *MapCounter) Inc(labels ...string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	k := c.key(labels...)
	c.values[k]++
}

// Value returns the count for the given label set (test helper).
func (c *MapCounter) Value(labels ...string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.values[c.key(labels...)]
}

// MetricsMiddleware increments counter once per completed request.
func MetricsMiddleware(counter Counter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// FIX ME: call c.Next(), then counter.Inc with method and status labels
		return c.Next()
	}
}
