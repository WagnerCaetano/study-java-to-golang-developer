package goroutine

import "sync"

func ParallelSum(nums []int, numGoroutines int) int {
	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	if len(nums) == 0 {
		return 0
	}
	if numGoroutines <= 0 {
		numGoroutines = 1
	}
	if numGoroutines > len(nums) {
		numGoroutines = len(nums)
	}

	chunkSize := len(nums) / numGoroutines
	remainder := len(nums) % numGoroutines
	start := 0

	for i := 0; i < numGoroutines; i++ {
		size := chunkSize
		if i < remainder {
			size++
		}
		end := start + size
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
		start = end
	}

	wg.Wait()
	return total
}
