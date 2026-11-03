package metrics

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Counter interface {
	Inc(labels ...string)
}

type MapCounter struct {
	mu     sync.Mutex
	values map[string]int64
}

func NewMapCounter() *MapCounter {
	return &MapCounter{values: make(map[string]int64)}
}

func (c *MapCounter) key(labels ...string) string {
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

func (c *MapCounter) Value(labels ...string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.values[c.key(labels...)]
}

func MetricsMiddleware(counter Counter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := c.Next()
		counter.Inc("method", c.Method(), "status", fmt.Sprintf("%d", c.Response().StatusCode()))
		return err
	}
}
