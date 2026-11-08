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

func TestParallelSumEdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		nums          []int
		numGoroutines int
		want          int
	}{
		{name: "empty input", nums: nil, numGoroutines: 4, want: 0},
		{name: "more goroutines than numbers", nums: []int{1, 2, 3}, numGoroutines: 5, want: 6},
		{name: "zero goroutines", nums: []int{1, 2, 3}, numGoroutines: 0, want: 6},
		{name: "negative goroutines", nums: []int{1, 2, 3}, numGoroutines: -2, want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParallelSum(tt.nums, tt.numGoroutines); got != tt.want {
				t.Errorf("ParallelSum = %d, want %d", got, tt.want)
			}
		})
	}
}
