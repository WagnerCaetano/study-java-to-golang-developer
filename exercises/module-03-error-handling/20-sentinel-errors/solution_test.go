package sentinel

import (
	"errors"
	"testing"
)

func TestSentinelMessages(t *testing.T) {
	if ErrNotFound.Error() != "user not found" {
		t.Errorf("ErrNotFound = %q", ErrNotFound.Error())
	}
	if ErrAlreadyExists.Error() != "user already exists" {
		t.Errorf("ErrAlreadyExists = %q", ErrAlreadyExists.Error())
	}
	if ErrInvalidInput.Error() != "invalid input" {
		t.Errorf("ErrInvalidInput = %q", ErrInvalidInput.Error())
	}
}

func TestFindUser(t *testing.T) {
	if !errors.Is(FindUser(-1), ErrInvalidInput) {
		t.Error("negative id should return ErrInvalidInput")
	}
	if !errors.Is(FindUser(999), ErrNotFound) {
		t.Error("id 999 should return ErrNotFound")
	}
	if FindUser(1) != nil {
		t.Error("valid id should return nil")
	}
}
