package circuitbreaker

import (
	"errors"
	"testing"
	"time"
)

func TestCircuitBreakerClosed(t *testing.T) {
	cb := New(3, 1*time.Second)
	err := cb.Execute(func() error { return nil })
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cb.State() != Closed {
		t.Error("should stay closed on success")
	}
}

func TestCircuitBreakerOpens(t *testing.T) {
	cb := New(3, 1*time.Second)
	for i := 0; i < 3; i++ {
		cb.Execute(func() error { return errors.New("fail") })
	}
	if cb.State() != Open {
		t.Error("should be open after threshold failures")
	}
}
