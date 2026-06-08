package waitgroup

import "sync"

// RunAll runs all functions concurrently and waits for completion
// FIX ME: Add wg.Add(1) before goroutine and wg.Done() inside
func RunAll(fns []func()) {
	var wg sync.WaitGroup
	for _, fn := range fns {
		go func(f func()) {
			// FIX: Add wg.Add(1) and wg.Done()
			f()
		}(fn)
	}
	wg.Wait()
}
