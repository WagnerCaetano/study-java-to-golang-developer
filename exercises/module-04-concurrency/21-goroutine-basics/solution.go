package goroutine

// ParallelSum sums a slice using numGoroutines goroutines.
func ParallelSum(nums []int, numGoroutines int) int {
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
	results := make(chan int, numGoroutines)
	start := 0

	for i := 0; i < numGoroutines; i++ {
		size := chunkSize
		if i < remainder {
			size++
		}
		end := start + size
		go func(slice []int) {
			sum := 0
			for _, n := range slice {
				sum += n
			}
			results <- sum
		}(nums[start:end])
		start = end
	}

	return <-results // TODO: collect and add the result from every goroutine
}
