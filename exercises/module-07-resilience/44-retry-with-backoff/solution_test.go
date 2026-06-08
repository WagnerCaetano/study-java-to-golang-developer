package retry

import (
	"errors"
	"testing"
)

func TestRetrySuccess(t *testing.T) {
	callCount := 0
	err := Retry(3, 0, func() error {
		callCount++
		if callCount < 2 {
			return errors.New("not yet")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if callCount != 2 {
		t.Errorf("callCount = %d, want 2", callCount)
	}
}

func TestRetryMaxAttempts(t *testing.T) {
	err := Retry(3, 0, func() error {
		return errors.New("always fail")
	})
	if err == nil {
		t.Error("expected error after max retries")
	}
}
