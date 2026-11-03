# Exercise 47: graceful-shutdown

## Problem

Simulate **graceful shutdown** of a server without binding a real listener.

- Background jobs run until their context is cancelled.
- `Shutdown` cancels that context, waits for background work to finish (`sync.WaitGroup`), or returns early if the shutdown `ctx` is done.

## API

```go
type Server struct { ... }
func NewServer() *Server
func (s *Server) StartBackgroundWork(fn func(ctx context.Context))
func (s *Server) Shutdown(ctx context.Context) error
```

`StartBackgroundWork` should run `fn` in a goroutine. `fn` is expected to loop until `ctx` is cancelled.

## How to Test

```bash
cd exercises/module-07-resilience/47-graceful-shutdown
go test -v
```

To verify the reference solution:

```bash
cp _solution/solution.go solution.go && go test -v
```
