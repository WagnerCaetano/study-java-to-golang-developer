package fiberrouting

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetUser(t *testing.T) {
	app := fiber.New()
	SetupRoutes(app)
	req := httptest.NewRequest("GET", "/users/42", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
}
