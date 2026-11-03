package fiberjson

import "github.com/gofiber/fiber/v2"

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Post("/users", func(c *fiber.Ctx) error {
		var req CreateUserRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid body"})
		}
		return c.Status(fiber.StatusCreated).JSON(req)
	})
	return app
}
