package unittest

import "testing"

func TestStringReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Go", "oG"},
	}
	for _, tt := range tests {
		result := StringReverse(tt.input)
		if result != tt.expected {
			t.Errorf("StringReverse(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
