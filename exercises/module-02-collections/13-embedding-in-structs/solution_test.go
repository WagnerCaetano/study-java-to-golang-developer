package embed

import "testing"

func TestCarEmbedding(t *testing.T) {
	c := NewCar("Toyota", "Corolla", 2024, 4)
	if c == nil {
		t.Fatal("NewCar returned nil")
	}
	if c.Make != "Toyota" {
		t.Errorf("Make = %q, want %q", c.Make, "Toyota")
	}
	if c.Model != "Corolla" {
		t.Errorf("Model = %q, want %q", c.Model, "Corolla")
	}
	if c.Year != 2024 {
		t.Errorf("Year = %d, want %d", c.Year, 2024)
	}
	if c.Doors != 4 {
		t.Errorf("Doors = %d, want %d", c.Doors, 4)
	}
}
