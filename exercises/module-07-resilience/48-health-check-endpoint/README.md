# Exercise 48: health-check-endpoint

## Problem

Expose **liveness** and **readiness** endpoints on a Fiber app.

| Route     | When              | Status | JSON body                    |
|-----------|-------------------|--------|------------------------------|
| `GET /health` | always          | 200    | `{"status":"ok"}`            |
| `GET /ready`  | `Ready == true` | 200    | `{"status":"ready"}`         |
| `GET /ready`  | `Ready == false`| 503    | `{"status":"not_ready"}`     |

## API

```go
type Checker struct { Ready bool }
func SetupApp(c *Checker) *fiber.App
```

Use `c.JSON(...)` (or equivalent) for JSON responses.

## How to Test

```bash
cd exercises/module-07-resilience/48-health-check-endpoint
go test -v
```

To verify the reference solution:

```bash
cp _solution/solution.go solution.go && go test -v
```
