package slicecap

import "testing"

func TestMerge(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	result := Merge(a, b)
	expected := []int{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Fatalf("len = %d, want %d", len(result), len(expected))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("result[%d] = %d, want %d", i, result[i], expected[i])
		}
	}
}

func TestMergeEmpty(t *testing.T) {
	result := Merge([]int{}, []int{})
	if len(result) != 0 {
		t.Errorf("empty merge should return empty slice")
	}
}
