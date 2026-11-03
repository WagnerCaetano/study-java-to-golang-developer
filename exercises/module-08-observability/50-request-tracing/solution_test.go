package reqtrace

import (
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRequestIDGeneratedWhenMissing(t *testing.T) {
	app := fiber.New()
	app.Use(RequestIDMiddleware())
	app.Get("/", func(c *fiber.Ctx) error {
		id := GetRequestID(c)
		if id == "" {
			t.Error("expected request id in locals")
		}
		return c.SendString(id)
	})

	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	id := resp.Header.Get(HeaderRequestID)
	if id == "" {
		t.Fatal("missing response header")
	}
	if !regexp.MustCompile(`^[0-9a-f]{32}$`).MatchString(id) {
		t.Errorf("id %q is not 32-char hex", id)
	}
}

func TestRequestIDPreservedFromClient(t *testing.T) {
	const clientID = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	app := fiber.New()
	app.Use(RequestIDMiddleware())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(GetRequestID(c))
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(HeaderRequestID, clientID)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}
	if got := resp.Header.Get(HeaderRequestID); got != clientID {
		t.Errorf("header = %q, want %q", got, clientID)
	}
}
