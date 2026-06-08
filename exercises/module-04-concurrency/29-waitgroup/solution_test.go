package waitgroup

import (
	"sync/atomic"
	"testing"
)

func TestRunAll(t *testing.T) {
	var count int32
	fns := make([]func(), 10)
	for i := range fns {
		fns[i] = func() {
			atomic.AddInt32(&count, 1)
		}
	}
	RunAll(fns)
	if atomic.LoadInt32(&count) != 10 {
		t.Errorf("count = %d, want 10", count)
	}
}
