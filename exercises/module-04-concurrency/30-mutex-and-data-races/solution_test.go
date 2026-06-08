package mutex

import "testing"

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if c.Get() != 3 {
		t.Errorf("Get = %d, want 3", c.Get())
	}
	c.Reset()
	if c.Get() != 0 {
		t.Errorf("After reset: Get = %d, want 0", c.Get())
	}
}
