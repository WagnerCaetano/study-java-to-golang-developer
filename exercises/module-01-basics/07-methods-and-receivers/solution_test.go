package methods

import "testing"

func TestCounterIncrement(t *testing.T) {
	c := Counter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if c.Value() != 3 {
		t.Errorf("Value = %d, want 3", c.Value())
	}
}

func TestCounterReset(t *testing.T) {
	c := Counter{}
	c.Increment()
	c.Increment()
	c.Reset()
	if c.Value() != 0 {
		t.Errorf("After reset, Value = %d, want 0", c.Value())
	}
}
