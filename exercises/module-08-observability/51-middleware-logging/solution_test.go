package mwlog

import (
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestLoggingMiddlewareFields(t *testing.T) {
	var buf bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&buf, &slog.HandlerOptions{Level: slog.LevelInfo}))

	app := fiber.New()
	app.Use(LoggingMiddleware(logger))
	app.Get("/items", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/items", nil)
	if _, err := app.Test(req, -1); err != nil {
		t.Fatal(err)
	}

	line := strings.TrimSpace(buf.String())
	var m map[string]any
	if err := json.Unmarshal([]byte(line), &m); err != nil {
		t.Fatalf("log JSON: %v, raw %q", err, line)
	}
	if m["method"] != "GET" {
		t.Errorf("method = %v", m["method"])
	}
	if m["path"] != "/items" {
		t.Errorf("path = %v", m["path"])
	}
	if status, ok := m["status"].(float64); !ok || int(status) != 200 {
		t.Errorf("status = %v, want 200", m["status"])
	}
	if _, ok := m["duration_ms"].(float64); !ok {
		t.Errorf("expected duration_ms number, got %v", m["duration_ms"])
	}
}
