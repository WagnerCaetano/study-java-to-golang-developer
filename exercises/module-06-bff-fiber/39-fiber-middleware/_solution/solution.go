package fiberex

import "github.com/gofiber/fiber/v2"

// SetupApp creates a Fiber app with the required configuration
func SetupApp() *fiber.App {
	app := fiber.New()
	// TODO: Implement the exercise requirements
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	return app
}
