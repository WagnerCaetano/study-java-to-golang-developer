# Exercise 45: rate-limiting

## Problem

Implement a **token-bucket** rate limiter.

- `NewLimiter(capacity, refillEvery)` starts with `capacity` tokens available.
- Each successful `Allow()` consumes one token and returns `true`.
- When no tokens are available, `Allow()` returns `false`.
- On each `Allow()`, refill tokens based on **elapsed time**: add one token every `refillEvery`, capped at `capacity`.

## API

```go
type Limiter struct { ... }
func NewLimiter(capacity int, refillEvery time.Duration) *Limiter
func (l *Limiter) Allow() bool
```

## Hints

- Track `lastRefill` (or similar) and compute how many whole refill intervals passed since the last update.
- Use a `sync.Mutex` if multiple goroutines may call `Allow()`.

## How to Test

```bash
cd exercises/module-07-resilience/45-rate-limiting
go test -v
```

To verify the reference solution:

```bash
cp _solution/solution.go solution.go && go test -v
```
