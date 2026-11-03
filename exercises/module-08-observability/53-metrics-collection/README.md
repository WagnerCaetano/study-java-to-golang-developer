# Exercise 53: Metrics Collection

## Problem

HTTP services expose **request counters** for Prometheus-style monitoring. In Fiber apps, [`github.com/ansrivas/fiberprometheus/v2`](https://github.com/ansrivas/fiberprometheus) registers `/metrics` and middleware automatically. Here you implement the same idea against a tiny **`Counter` interface** so tests do not scrape an endpoint.

Package **`metrics`**:

1. `Counter` with `Inc(labels ...string)` — labels are key/value pairs (`"method", "GET", "status", "200"`).
2. `NewMapCounter() *MapCounter` — thread-safe in-memory counter; `Value(labels ...string) int64` for tests.
3. `MetricsMiddleware(counter Counter) fiber.Handler` — after the request, increment with method and stringified status code.

## Learning objectives

- Labelled counters (Prometheus cardinality basics).
- Middleware as the hook for RED metrics (rate, errors, duration — here: rate by status).
- Interface-first design before pulling in fiberprometheus.

## How to test

```bash
go test -v
```
