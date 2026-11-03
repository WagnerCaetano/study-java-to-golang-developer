package fibererr

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestBoomFiberError(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/boom", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusBadRequest)
	}
	assertErrorJSON(t, resp, "bad request")
}

func TestPanicIshGenericError(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/panic-ish", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("app.Test error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != fiber.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", resp.StatusCode, fiber.StatusInternalServerError)
	}
	assertErrorJSON(t, resp, "unexpected")
}

func assertErrorJSON(t *testing.T, resp *http.Response, wantMsg string) {
	t.Helper()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}
	var payload map[string]string
	if err := json.Unmarshal(body, &payload); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if payload["error"] != wantMsg {
		t.Errorf("error = %q, want %q", payload["error"], wantMsg)
	}
}
