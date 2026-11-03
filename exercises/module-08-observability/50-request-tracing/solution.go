package reqtrace

import "github.com/gofiber/fiber/v2"

const HeaderRequestID = "X-Request-ID"

const localRequestID = "request_id"

// GenerateRequestID returns a new unique request identifier.
func GenerateRequestID() string {
	// FIX ME: read 16 bytes from crypto/rand and hex-encode them
	return "00000000000000000000000000000000"
}

// RequestIDMiddleware ensures each request has X-Request-ID.
func RequestIDMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// FIX ME: reuse incoming header or generate; set locals and response header
		return c.Next()
	}
}

// GetRequestID returns the request ID stored by the middleware.
func GetRequestID(c *fiber.Ctx) string {
	// FIX ME: read from c.Locals(localRequestID)
	return ""
}
