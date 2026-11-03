package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

// Checker holds readiness state for dependency checks.
type Checker struct {
	Ready bool
}

// SetupApp registers /health and /ready routes.
// FIX ME: Implement liveness and readiness responses per README.
func SetupApp(c *Checker) *fiber.App {
	app := fiber.New()
	_ = c
	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ok")
	})
	app.Get("/ready", func(ctx *fiber.Ctx) error {
		return ctx.SendString("ready")
	})
	return app
}
