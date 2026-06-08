package buffered

import "sync"

// BoundedProcessor runs tasks with at most maxConcurrent goroutines
// FIX ME: Use a buffered channel as a semaphore
func BoundedProcessor(tasks []func(), maxConcurrent int) {
	var wg sync.WaitGroup
	// FIX: Create a buffered channel semaphore and use it

	for _, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
		}(task)
	}
	wg.Wait()
}
