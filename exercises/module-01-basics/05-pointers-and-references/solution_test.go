package pointers

import "testing"

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("Swap: a=%d, b=%d; want a=2, b=1", a, b)
	}
}

func TestSwapEqual(t *testing.T) {
	a, b := 5, 5
	Swap(&a, &b)
	if a != 5 || b != 5 {
		t.Errorf("Swap equal: a=%d, b=%d; want a=5, b=5", a, b)
	}
}
