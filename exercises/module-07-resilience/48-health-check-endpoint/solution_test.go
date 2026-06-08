package resilience

import "testing"

func TestPlaceholder(t *testing.T) {
	if !Placeholder() {
		t.Error("placeholder should return true")
	}
}
