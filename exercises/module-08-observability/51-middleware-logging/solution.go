package mwlog

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoggingMiddleware logs one structured line per request.
func LoggingMiddleware(logger *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// FIX ME: record start time, invoke c.Next(), log method/path/status/duration_ms
		return c.Next()
	}
}

// unused — prevents compile error until you implement timing
var _ = time.Now
