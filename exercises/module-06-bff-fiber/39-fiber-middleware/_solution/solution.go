package fibermw

import (
	"fmt"
	"sync/atomic"

	"github.com/gofiber/fiber/v2"
)

// RequestIDGenerator produces a request ID when the client did not send X-Request-ID.
type RequestIDGenerator func() string

var requestIDSeq atomic.Uint64

func defaultRequestIDGenerator() string {
	return fmt.Sprintf("req-%d", requestIDSeq.Add(1))
}

// RequestIDMiddleware ensures responses include X-Request-ID.
func RequestIDMiddleware(generate RequestIDGenerator) fiber.Handler {
	if generate == nil {
		generate = defaultRequestIDGenerator
	}
	return func(c *fiber.Ctx) error {
		rid := c.Get("X-Request-ID")
		if rid == "" {
			rid = generate()
		}
		c.Set("X-Request-ID", rid)
		return c.Next()
	}
}

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Use(RequestIDMiddleware(nil))
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	return app
}
