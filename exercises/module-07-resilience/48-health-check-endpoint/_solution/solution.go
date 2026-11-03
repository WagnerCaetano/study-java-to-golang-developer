package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

type Checker struct {
	Ready bool
}

func SetupApp(c *Checker) *fiber.App {
	app := fiber.New()

	app.Get("/health", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
		})
	})

	app.Get("/ready", func(ctx *fiber.Ctx) error {
		if c.Ready {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"status": "ready",
			})
		}
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"status": "not_ready",
		})
	})

	return app
}
