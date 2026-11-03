package healthcheck

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func readStatus(t *testing.T, resp *http.Response) string {
	t.Helper()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var payload map[string]string
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("json unmarshal %q: %v", body, err)
	}
	return payload["status"]
}

func TestHealthAlwaysOK(t *testing.T) {
	app := SetupApp(&Checker{Ready: false})
	req := httptest.NewRequest("GET", "/health", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
	if got := readStatus(t, resp); got != "ok" {
		t.Fatalf("status field = %q, want ok", got)
	}
}

func TestReadyWhenReady(t *testing.T) {
	app := SetupApp(&Checker{Ready: true})
	req := httptest.NewRequest("GET", "/ready", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
	if got := readStatus(t, resp); got != "ready" {
		t.Fatalf("status field = %q, want ready", got)
	}
}

func TestReadyWhenNotReady(t *testing.T) {
	app := SetupApp(&Checker{Ready: false})
	req := httptest.NewRequest("GET", "/ready", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test: %v", err)
	}
	if resp.StatusCode != fiber.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusServiceUnavailable)
	}
	if got := readStatus(t, resp); got != "not_ready" {
		t.Fatalf("status field = %q, want not_ready", got)
	}
}
