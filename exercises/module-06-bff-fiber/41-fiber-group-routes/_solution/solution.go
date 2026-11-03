package fibergroup

import "github.com/gofiber/fiber/v2"

func SetupApp() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/api/v1")
	users := v1.Group("/users")
	users.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"users": []any{}})
	})
	users.Get("/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"id": c.Params("id")})
	})
	return app
}
