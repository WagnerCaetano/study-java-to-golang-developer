package fibererr

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

// SetupApp creates a Fiber app with a custom ErrorHandler and demo routes.
// FIX ME: JSON error handler; GET /boom and GET /panic-ish as described in README.
func SetupApp() *fiber.App {
	app := fiber.New()
	app.Get("/boom", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "bad request")
	})
	app.Get("/panic-ish", func(c *fiber.Ctx) error {
		return errors.New("unexpected")
	})
	return app
}
