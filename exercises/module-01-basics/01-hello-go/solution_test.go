package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	expected := "Hello, Go!"
	if result != expected {
		t.Errorf("Greet() = %q, want %q", result, expected)
	}
}
