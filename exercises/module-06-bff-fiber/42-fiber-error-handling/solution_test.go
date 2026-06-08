package fiberex

import (
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
)

func TestPing(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
}
