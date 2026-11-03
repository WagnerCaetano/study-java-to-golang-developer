# Exercise 57: Parallel HTTP Requests

## Problem

BFFs often fan out to several upstream HTTP APIs. Implement concurrent fetches behind an interface so tests never touch the real network.

Package **`parallelhttp`**:

```go
type Fetcher interface {
    Fetch(ctx context.Context, url string) (string, error)
}

func FetchAll(ctx context.Context, f Fetcher, urls []string) ([]string, error)
```

Rules:

- Fetch every URL **concurrently** (use `errgroup` or `WaitGroup` + mutex).
- **Preserve input order** in the returned slice (index `i` matches `urls[i]`).
- On the **first error**, cancel remaining work and return that error (derived context cancellation is recommended).
- Respect `ctx` cancellation.

## Learning objectives

- Fan-out with bounded concurrency patterns.
- Ordered aggregation from parallel goroutines.
- Fail-fast and context propagation in a BFF-style fetch layer.

## How to test

```bash
go test -v
```
