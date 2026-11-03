package ratelimit

import (
	"sync"
	"time"
)

type Limiter struct {
	mu          sync.Mutex
	capacity    int
	tokens      int
	refillEvery time.Duration
	lastRefill  time.Time
}

func NewLimiter(capacity int, refillEvery time.Duration) *Limiter {
	return &Limiter{
		capacity:    capacity,
		tokens:      capacity,
		refillEvery: refillEvery,
		lastRefill:  time.Now(),
	}
}

func (l *Limiter) refillLocked(now time.Time) {
	if l.refillEvery <= 0 {
		return
	}
	elapsed := now.Sub(l.lastRefill)
	add := int(elapsed / l.refillEvery)
	if add <= 0 {
		return
	}
	l.tokens += add
	if l.tokens > l.capacity {
		l.tokens = l.capacity
	}
	l.lastRefill = l.lastRefill.Add(time.Duration(add) * l.refillEvery)
}

func (l *Limiter) Allow() bool {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.refillLocked(time.Now())
	if l.tokens <= 0 {
		return false
	}
	l.tokens--
	return true
}
