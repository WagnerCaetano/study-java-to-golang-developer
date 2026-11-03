# Exercise 61: Redis Get / Set

## Problem

Caching layers sit behind a small API. Production Redis clients use [`github.com/redis/go-redis/v9`](https://github.com/redis/go-redis) — this exercise uses a **`Client` interface** so tests run offline.

Package **`redisex`**:

1. `Client` — `Get(ctx, key) (string, error)`, `Set(ctx, key, val string, ttl time.Duration) error`.
2. `MemoryClient` — in-memory implementation (provided in tests).
3. `CacheStore` — wraps `Client` with `Get` / `Set` delegating to the client.
4. `NewCacheStore(c Client) *CacheStore`.

`Get` returns `ErrCacheMiss` when the key is absent (wrap memory client behavior).

## Learning objectives

- Interface-driven cache boundaries (swap Redis for memory in tests).
- TTL semantics (expiration checks on read).
- Wiring go-redis in real apps: `redis.NewClient(&redis.Options{Addr: ...})`.

## How to test

```bash
go test -v
```
