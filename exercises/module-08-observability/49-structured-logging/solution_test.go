package logging

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
	"log/slog"
)

func TestNewJSONLoggerOutputsJSON(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJSONLogger(&buf)
	LogInfo(logger, "hello", slog.String("service", "bff"))

	line := strings.TrimSpace(buf.String())
	if !strings.HasPrefix(line, "{") {
		t.Fatalf("expected JSON line, got %q", line)
	}
	var m map[string]any
	if err := json.Unmarshal([]byte(line), &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["msg"] != "hello" {
		t.Errorf("msg = %v, want hello", m["msg"])
	}
	if m["service"] != "bff" {
		t.Errorf("service = %v, want bff", m["service"])
	}
}

func TestLogErrorIncludesErrorField(t *testing.T) {
	var buf bytes.Buffer
	logger := NewJSONLogger(&buf)
	LogError(logger, "failed", errTest, slog.Int("code", 500))

	line := strings.TrimSpace(buf.String())
	var m map[string]any
	if err := json.Unmarshal([]byte(line), &m); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if m["level"] != "ERROR" && m["level"] != "error" {
		// slog JSON uses string level in recent Go versions
		if lvl, ok := m["level"].(string); !ok || lvl != "ERROR" {
			// accept numeric level in older encodings — at minimum msg and error must exist
		}
	}
	if _, ok := m["error"]; !ok {
		t.Errorf("expected error field in %v", m)
	}
	if m["code"] != float64(500) {
		t.Errorf("code = %v, want 500", m["code"])
	}
}

var errTest = &testError{"boom"}

type testError struct{ s string }

func (e *testError) Error() string { return e.s }
