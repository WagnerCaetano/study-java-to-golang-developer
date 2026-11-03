package ratelimit

import "time"

// Limiter is a token-bucket rate limiter.
type Limiter struct {
	capacity    int
	tokens      int
	refillEvery time.Duration
	lastRefill  time.Time
}

// NewLimiter creates a limiter with the given capacity and refill interval.
func NewLimiter(capacity int, refillEvery time.Duration) *Limiter {
	return &Limiter{
		capacity:    capacity,
		tokens:      capacity,
		refillEvery: refillEvery,
		lastRefill:  time.Now(),
	}
}

// Allow reports whether a request is permitted and consumes one token when true.
// FIX ME: Refill tokens based on elapsed time before deciding.
func (l *Limiter) Allow() bool {
	if l.tokens > 0 {
		l.tokens--
		return true
	}
	return false
}
