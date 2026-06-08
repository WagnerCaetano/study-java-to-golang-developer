package fanout

import "sync"

func ProcessFanOut(items []int, numWorkers int, process func(int) int) []int {
	results := make([]int, len(items))
	var wg sync.WaitGroup
	chunkSize := (len(items) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(items) {
			end = len(items)
		}
		if start >= len(items) {
			break
		}
		wg.Add(1)
		go func(idx int, slice []int) {
			defer wg.Done()
			for j, item := range slice {
				results[idx+j] = process(item)
			}
		}(start, items[start:end])
	}
	wg.Wait()
	return results
}
