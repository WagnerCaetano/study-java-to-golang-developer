package metrics

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestMetricsMiddlewareIncrements(t *testing.T) {
	counter := NewMapCounter()
	app := fiber.New()
	app.Use(MetricsMiddleware(counter))
	app.Get("/ok", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	req := httptest.NewRequest("GET", "/ok", nil)
	if _, err := app.Test(req, -1); err != nil {
		t.Fatal(err)
	}

	n := counter.Value("method", "GET", "status", "200")
	if n != 1 {
		t.Errorf("count = %d, want 1", n)
	}
}

func TestMapCounterLabels(t *testing.T) {
	c := NewMapCounter()
	c.Inc("method", "POST", "status", "201")
	c.Inc("method", "POST", "status", "201")
	if v := c.Value("method", "POST", "status", "201"); v != 2 {
		t.Errorf("value = %d, want 2", v)
	}
}
