package circuitbreaker

import "time"

type State int

const (
	Closed   State = iota
	Open
	HalfOpen
)

// CircuitBreaker prevents cascading failures
// FIX ME: Implement the circuit breaker logic
type CircuitBreaker struct {
	failureCount    int
	failureThreshold int
	state           State
	lastFailure     time.Time
	timeout         time.Duration
}

func New(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		timeout:         timeout,
		state:           Closed,
	}
}

// Execute runs fn if the circuit is closed or half-open
// FIX ME: Track failures and change state
func (cb *CircuitBreaker) Execute(fn func() error) error {
	// FIX: Check state, execute if allowed, track failures
	return nil
}

func (cb *CircuitBreaker) State() State {
	return cb.state
}
