package buffered

import "sync"

func BoundedProcessor(tasks []func(), maxConcurrent int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent)

	for _, task := range tasks {
		sem <- struct{}{}
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
			<-sem
		}(task)
	}
	wg.Wait()
}
