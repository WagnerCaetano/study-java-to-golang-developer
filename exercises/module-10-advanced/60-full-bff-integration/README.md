# Exercise 60: Full BFF Integration

## Problem

Wire aggregation, caching, and HTTP into a small Fiber BFF — the capstone for this module.

Package **`bffinteg`**:

```go
type Deps struct {
    Users  UserClient  // GetUser
    Orders OrderClient // ListOrders
    Cache  interface {
        Get(k string) (string, bool)
        Set(k, v string)
    }
}

func SetupApp(d Deps) *fiber.App
```

Routes:

- **`GET /health`** → body `ok` with status 200.
- **`GET /dashboard/:userID`**
  - Cache key: `dash:{userID}`.
  - **Hit:** respond with the cached JSON string, `Content-Type: application/json`.
  - **Miss:** build dashboard (user + orders **in parallel**), JSON-encode, store in cache, return 200.
  - **Upstream error:** status **502** with body `{"error":"<message>"}`.

Define `User`, `Order`, `Dashboard`, and client interfaces in this package (no imports from other exercises).

## Learning objectives

- End-to-end BFF handler design.
- Cache-aside for aggregated responses.
- Mapping backend failures to gateway-style HTTP errors.

## How to test

```bash
go test -v
```
