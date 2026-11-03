package fibermw

import "github.com/gofiber/fiber/v2"

// RequestIDGenerator produces a request ID when the client did not send X-Request-ID.
type RequestIDGenerator func() string

// RequestIDMiddleware ensures responses include X-Request-ID.
// FIX ME: Set the header when missing (use generate), echo client value when present, then c.Next().
func RequestIDMiddleware(generate RequestIDGenerator) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Next()
	}
}

// SetupApp creates a Fiber app with request ID middleware and GET /ping.
// FIX ME: Apply RequestIDMiddleware and register GET /ping → "pong".
func SetupApp() *fiber.App {
	app := fiber.New()
	return app
}
