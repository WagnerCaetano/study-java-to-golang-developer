package fibererr

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func jsonErrorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	var fe *fiber.Error
	if errors.As(err, &fe) {
		code = fe.Code
	}
	return c.Status(code).JSON(fiber.Map{"error": err.Error()})
}

func SetupApp() *fiber.App {
	app := fiber.New(fiber.Config{ErrorHandler: jsonErrorHandler})
	app.Get("/boom", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "bad request")
	})
	app.Get("/panic-ish", func(c *fiber.Ctx) error {
		return errors.New("unexpected")
	})
	return app
}
