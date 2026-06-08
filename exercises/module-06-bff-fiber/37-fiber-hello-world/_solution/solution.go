package fiberhello

import "github.com/gofiber/fiber/v2"

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	return app
}
