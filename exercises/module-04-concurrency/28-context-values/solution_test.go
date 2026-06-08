package ctxvalues

import (
	"context"
	"testing"
)

func TestRequestID(t *testing.T) {
	ctx := context.Background()
	ctx = WithRequestID(ctx, "abc-123")
	id := GetRequestID(ctx)
	if id != "abc-123" {
		t.Errorf("got %q, want %q", id, "abc-123")
	}
}

func TestMissingRequestID(t *testing.T) {
	ctx := context.Background()
	id := GetRequestID(ctx)
	if id != "" {
		t.Errorf("got %q, want empty string", id)
	}
}
