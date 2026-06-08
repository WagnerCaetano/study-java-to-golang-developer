package circuitbreaker

import (
	"errors"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	failureCount     int
	failureThreshold int
	state            State
	lastFailure      time.Time
	timeout          time.Duration
}

func New(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		timeout:         timeout,
		state:           Closed,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	if cb.state == Open {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = HalfOpen
		} else {
			return errors.New("circuit breaker is open")
		}
	}

	err := fn()
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		if cb.failureCount >= cb.failureThreshold {
			cb.state = Open
		}
		return err
	}

	cb.failureCount = 0
	cb.state = Closed
	return nil
}

func (cb *CircuitBreaker) State() State {
	return cb.state
}
