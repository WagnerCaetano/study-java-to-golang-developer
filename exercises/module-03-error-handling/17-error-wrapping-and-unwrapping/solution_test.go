package wrapping

import (
	"errors"
	"testing"
)

func TestReadConfigWrap(t *testing.T) {
	err := ReadConfig("/etc/app/config.yaml")
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, ErrConfigNotFound) {
		t.Errorf("error should wrap ErrConfigNotFound")
	}
}

func TestGetRootCause(t *testing.T) {
	err := ReadConfig("/test")
	root := GetRootCause(err)
	if !errors.Is(root, ErrConfigNotFound) {
		t.Errorf("root cause should be ErrConfigNotFound, got %v", root)
	}
}
