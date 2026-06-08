package goroutine

import "testing"

func TestParallelSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := ParallelSum(nums, 4)
	expected := 55
	if result != expected {
		t.Errorf("ParallelSum = %d, want %d", result, expected)
	}
}

func TestParallelSumSingle(t *testing.T) {
	nums := []int{1, 2, 3}
	result := ParallelSum(nums, 1)
	if result != 6 {
		t.Errorf("ParallelSum = %d, want 6", result)
	}
}
