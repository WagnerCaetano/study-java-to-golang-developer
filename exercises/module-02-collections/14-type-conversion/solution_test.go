package conversion

import "testing"

func TestCToF(t *testing.T) {
	tests := []struct {
		input    Celsius
		expected Fahrenheit
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
	}
	for _, tt := range tests {
		result := CToF(tt.input)
		if result != tt.expected {
			t.Errorf("CToF(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		input    Fahrenheit
		expected Celsius
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
	}
	for _, tt := range tests {
		result := FToC(tt.input)
		if result != tt.expected {
			t.Errorf("FToC(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
