package goroutine

import "sync"

func ParallelSum(nums []int, numGoroutines int) int {
	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunkSize := len(nums) / numGoroutines
	if chunkSize == 0 {
		chunkSize = len(nums)
		numGoroutines = 1
	}

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(nums)
		}
		if start >= len(nums) {
			break
		}
		wg.Add(1)
		go func(slice []int) {
			defer wg.Done()
			sum := 0
			for _, n := range slice {
				sum += n
			}
			mu.Lock()
			total += sum
			mu.Unlock()
		}(nums[start:end])
	}

	wg.Wait()
	return total
}
