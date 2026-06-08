package buffered

import (
	"sync/atomic"
	"testing"
)

func TestBoundedProcessor(t *testing.T) {
	var concurrent, maxConcurrent int32
	tasks := make([]func(), 20)
	for i := range tasks {
		tasks[i] = func() {
			c := atomic.AddInt32(&concurrent, 1)
			for {
				old := atomic.LoadInt32(&maxConcurrent)
				if c <= old || atomic.CompareAndSwapInt32(&maxConcurrent, old, c) {
					break
				}
			}
			atomic.AddInt32(&concurrent, -1)
		}
	}
	BoundedProcessor(tasks, 3)
	if atomic.LoadInt32(&maxConcurrent) > 3 {
		t.Errorf("max concurrent = %d, should be <= 3", maxConcurrent)
	}
}
