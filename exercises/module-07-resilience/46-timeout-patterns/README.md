# Exercise 46: timeout-patterns

## Problem

Implement `RunWithTimeout`, which runs `work` with a **deadline** derived from `d`.

- Create a child context with `context.WithTimeout`.
- Invoke `work` with that child context.
- If `work` respects cancellation and the deadline passes first, return the context error (typically `context.DeadlineExceeded`).
- If `work` finishes before the deadline, return `work`'s error (including `nil` on success).

Run `work` in a way that does not block forever when the deadline expires (e.g. goroutine + `select` on `ctx.Done()` and a result channel).

## API

```go
func RunWithTimeout(ctx context.Context, d time.Duration, work func(context.Context) error) error
```

## How to Test

```bash
cd exercises/module-07-resilience/46-timeout-patterns
go test -v
```

To verify the reference solution:

```bash
cp _solution/solution.go solution.go && go test -v
```
