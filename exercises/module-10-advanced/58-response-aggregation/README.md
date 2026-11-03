# Exercise 58: Response Aggregation

## Problem

A dashboard BFF loads user profile and order history from two services. Aggregate them in one call while minimizing latency.

Package **`aggregate`**:

```go
type User struct { ID, Name string }
type Order struct { ID, UserID string }
type Dashboard struct { User User; Orders []Order }

type UserClient interface {
    GetUser(ctx context.Context, id string) (User, error)
}
type OrderClient interface {
    ListOrders(ctx context.Context, userID string) ([]Order, error)
}

func BuildDashboard(ctx context.Context, userID string, users UserClient, orders OrderClient) (Dashboard, error)
```

Rules:

- Call **both** upstream clients **in parallel** (orders only need `userID`, which you already have).
- On error from either client, return a zero `Dashboard` and the error.
- On success, return `Dashboard{User: ..., Orders: ...}`.

## Learning objectives

- BFF aggregation without N+1 sequential calls.
- Composing multiple interface-backed clients.
- Testing with fakes instead of real HTTP.

## How to test

```bash
go test -v
```
