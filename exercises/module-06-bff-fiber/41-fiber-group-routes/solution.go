package fibergroup

import "github.com/gofiber/fiber/v2"

// SetupApp registers versioned API routes under /api/v1.
// FIX ME: Group /api/v1, subgroup /users with GET / and GET /:id.
func SetupApp() *fiber.App {
	app := fiber.New()
	return app
}
