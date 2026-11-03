# Exercise 62: Cache-Aside (GetOrLoad)

## Problem

**Cache-aside** loads data on miss, stores it, and returns the value. Implement `GetOrLoad` using the `Client` from Exercise 61.

Package **`cacheaside`**:

```go
func GetOrLoad(ctx context.Context, client redisex.Client, key string, ttl time.Duration, loader func(context.Context) (string, error)) (string, error)
```

On hit, return cached value. On `ErrCacheMiss`, call `loader`, `Set` with `ttl`, return value. Propagate loader errors without caching.

Import `redisex` from this module path for `Client` and `ErrCacheMiss`.

## Learning objectives

- Cache-aside vs read-through / write-through.
- Stampede mitigation (out of scope here — single loader call on miss).
- Composing generic helpers over interfaces.

## How to test

```bash
go test -v
```
