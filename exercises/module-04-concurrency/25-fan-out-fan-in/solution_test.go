package fanout

import "testing"

func TestFanOutFanIn(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := ProcessFanOut(items, 4, func(n int) int {
		return n * n
	})
	if len(result) != len(items) {
		t.Fatalf("len = %d, want %d", len(result), len(items))
	}
}
