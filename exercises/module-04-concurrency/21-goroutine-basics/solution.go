package goroutine

import "sync"

// ParallelSum sums a slice using numGoroutines goroutines
// FIX ME: Split the work and use goroutines + WaitGroup
func ParallelSum(nums []int, numGoroutines int) int {
	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunkSize := len(nums) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(nums)
		}
		wg.Add(1)
		go func(slice []int) {
			defer wg.Done()
			sum := 0
			for _, n := range slice {
				sum += n
			}
			mu.Lock()
			total += sum // FIX: This has a race condition! Use mutex properly.
			mu.Unlock()
		}(nums[start:end])
	}

	wg.Wait()
	return total
}
