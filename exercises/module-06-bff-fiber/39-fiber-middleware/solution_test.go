package fibermw

import (
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestPingSetsRequestID(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
	rid := resp.Header.Get("X-Request-ID")
	if rid == "" {
		t.Fatal("X-Request-ID header missing from response")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	if string(body) != "pong" {
		t.Errorf("body = %q, want %q", body, "pong")
	}
}

func TestPingEchoesClientRequestID(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/ping", nil)
	req.Header.Set("X-Request-ID", "abc")
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test error: %v", err)
	}
	defer resp.Body.Close()

	if got := resp.Header.Get("X-Request-ID"); got != "abc" {
		t.Errorf("X-Request-ID = %q, want %q", got, "abc")
	}
}
