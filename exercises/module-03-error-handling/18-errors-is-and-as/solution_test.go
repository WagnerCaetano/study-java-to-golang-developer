package errorsis

import (
	"fmt"
	"testing"
)

func TestCheckNotFound(t *testing.T) {
	err := fmt.Errorf("wrap: %w", &NotFoundError{Resource: "user", ID: 42})
	if !CheckNotFound(err) {
		t.Error("should detect NotFoundError in chain")
	}
	if CheckNotFound(fmt.Errorf("other error")) {
		t.Error("should not detect NotFoundError for other errors")
	}
}

func TestCheckUnauthorized(t *testing.T) {
	err := fmt.Errorf("wrap: %w", ErrUnauthorized)
	if !CheckUnauthorized(err) {
		t.Error("should detect ErrUnauthorized in chain")
	}
	if CheckUnauthorized(fmt.Errorf("other error")) {
		t.Error("should not detect ErrUnauthorized for other errors")
	}
}
