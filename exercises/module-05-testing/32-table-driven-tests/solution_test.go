package tabletest

import "testing"

func TestCalculator(t *testing.T) {
	c := Calculator{}
	t.Run("Add", func(t *testing.T) {
		tests := []struct{ a, b, expected int }{
			{1, 2, 3}, {0, 0, 0}, {-1, 1, 0},
		}
		for _, tt := range tests {
			if got := c.Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%d,%d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		}
	})
	t.Run("Divide by zero", func(t *testing.T) {
		_, err := c.Divide(1, 0)
		if err == nil {
			t.Error("expected error for division by zero")
		}
	})
}
