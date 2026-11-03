# Chapter 14: Redis with go-redis

> Redis is the default choice for BFF caching, session storage, and rate-limit counters. `redis/go-redis/v9` is the maintained client — context-first, pipelining, and cluster-aware.

---

## Java vs Go Redis Clients

| Capability | Java | Go |
|------------|------|-----|
| Sync API | Jedis | go-redis commands return directly |
| Async/reactive | Lettuce (Netty) | Goroutines + context; no separate reactive API |
| Connection pool | Jedis pool / Lettuce shared connection | Built into `redis.Options` |
| TTL | `SETEX`, `expire` | `Set` with `Expiration`, `Expire` |
| Pipelines | `Pipeline` | `Pipelined` / `Pipeline()` |
| Context/cancel | Less central in older APIs | Every command takes `context.Context` |

Spring Cache abstracts Redis behind `@Cacheable`. In Go BFFs you usually implement **cache-aside** explicitly — clearer latency and failure behavior.

---

## Setup

```bash
go get github.com/redis/go-redis/v9
```

Redis 7+ locally:

```bash
docker run -d --name redis -p 6379:6379 redis:7-alpine
```

---

## Client Configuration

```go
package cache

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
)

func NewRedis(addr, password string, db int) *redis.Client {
    return redis.NewClient(&redis.Options{
        Addr:         addr, // "localhost:6379"
        Password:     password,
        DB:           db,
        DialTimeout:  5 * time.Second,
        ReadTimeout:  3 * time.Second,
        WriteTimeout: 3 * time.Second,
        PoolSize:     20,
        MinIdleConns: 5,
    })
}

func Ping(ctx context.Context, rdb *redis.Client) error {
    ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
    defer cancel()
    return rdb.Ping(ctx).Err()
}
```

**Java comparison (Lettuce):**

```java
RedisClient client = RedisClient.create("redis://localhost:6379");
StatefulRedisConnection<String, String> conn = client.connect();
conn.sync().set("key", "value");
```

Go keeps a long-lived `*redis.Client` (pool inside) — create once in `main`, inject into handlers.

---

## Get, Set, Del, TTL

```go
func CacheUser(ctx context.Context, rdb *redis.Client, id, json string) error {
    key := "user:" + id
    return rdb.Set(ctx, key, json, 15*time.Minute).Err()
}

func GetCachedUser(ctx context.Context, rdb *redis.Client, id string) (string, error) {
    key := "user:" + id
    val, err := rdb.Get(ctx, key).Result()
    if err == redis.Nil {
        return "", ErrCacheMiss
    }
    return val, err
}

func InvalidateUser(ctx context.Context, rdb *redis.Client, id string) error {
    return rdb.Del(ctx, "user:"+id).Err()
}

func TouchTTL(ctx context.Context, rdb *redis.Client, id string) error {
    return rdb.Expire(ctx, "user:"+id, 15*time.Minute).Err()
}
```

`redis.Nil` means key missing — map it to your domain `ErrCacheMiss`, not a 500.

---

## Context Timeouts (Every Call)

BFF handlers should bound Redis latency:

```go
func withRedisTimeout(parent context.Context) (context.Context, context.CancelFunc) {
    return context.WithTimeout(parent, 200*time.Millisecond)
}

func (h *Handler) getUser(c *fiber.Ctx) error {
    ctx, cancel := withRedisTimeout(c.UserContext())
    defer cancel()

    id := c.Params("id")
    cached, err := GetCachedUser(ctx, h.rdb, id)
    if err == nil {
        return c.Type("json").SendString(cached)
    }
    if err != ErrCacheMiss && err != context.DeadlineExceeded {
        h.log.Warn("redis get failed", "error", err)
        // Often continue to origin — cache is optional
    }
    // ... fetch from origin, populate cache
    return nil
}
```

This matches resilience patterns from Chapter 11 — cache failure must not take down the BFF.

---

## Pipelines (Batch Round-Trips)

When aggregating a dashboard, you may need many keys:

```go
func MGetUsers(ctx context.Context, rdb *redis.Client, ids []string) ([]*string, error) {
    pipe := rdb.Pipeline()
    cmds := make([]*redis.StringCmd, len(ids))
    for i, id := range ids {
        cmds[i] = pipe.Get(ctx, "user:"+id)
    }
    _, err := pipe.Exec(ctx)
    if err != nil && err != redis.Nil {
        return nil, err
    }

    out := make([]*string, len(ids))
    for i, cmd := range cmds {
        s, err := cmd.Result()
        if err == redis.Nil {
            continue
        }
        if err != nil {
            return nil, err
        }
        out[i] = &s
    }
    return out, nil
}
```

**Java:** `connection.openPipeline()` / `sync()` — same idea: one network flush, multiple commands.

---

## Cache-Aside Pattern for a BFF

```go
type UserService struct {
    rdb    *redis.Client
    origin OriginClient // HTTP/gRPC to user service
}

func (s *UserService) GetUser(ctx context.Context, id string) ([]byte, error) {
    key := "bff:user:" + id

    if b, err := s.rdb.Get(ctx, key).Bytes(); err == nil {
        return b, nil
    } else if err != redis.Nil {
        // Log and fall through — stale cache optional
        slog.Warn("redis get", "key", key, "error", err)
    }

    body, err := s.origin.FetchUser(ctx, id)
    if err != nil {
        return nil, err
    }

    // Best-effort cache populate; don't fail request if SET fails
    if err := s.rdb.Set(ctx, key, body, 10*time.Minute).Err(); err != nil {
        slog.Warn("redis set", "key", key, "error", err)
    }
    return body, nil
}
```

**Stampede protection:** for hot keys, add singleflight (`golang.org/x/sync/singleflight`) around origin fetches — not shown here but common in production.

**Invalidation:** on webhook/event, call `Del` or publish to a channel; BFFs often use short TTLs instead of perfect invalidation.

---

## JSON and Structured Values

Store JSON as string or use RedisJSON module if your platform provides it. Simple approach:

```go
type Session struct {
    UserID string `json:"user_id"`
    Roles  []string `json:"roles"`
}

func SetSession(ctx context.Context, rdb *redis.Client, token string, s Session, ttl time.Duration) error {
    b, err := json.Marshal(s)
    if err != nil {
        return err
    }
    return rdb.Set(ctx, "session:"+token, b, ttl).Err()
}
```

---

## Health Checks

Include Redis in readiness when the BFF depends on cache for overload protection:

```go
func RedisHealth(rdb *redis.Client) fiber.Handler {
    return func(c *fiber.Ctx) error {
        ctx, cancel := context.WithTimeout(c.UserContext(), 500*time.Millisecond)
        defer cancel()
        if err := rdb.Ping(ctx).Err(); err != nil {
            return c.Status(503).JSON(fiber.Map{"redis": "down"})
        }
        return c.JSON(fiber.Map{"redis": "up"})
    }
}
```

If cache is **best-effort**, liveness can stay green while readiness still serves without Redis.

---

## Testing

Use `github.com/alicebob/miniredis/v2` or Docker in integration tests:

```go
func TestCacheAside(t *testing.T) {
    mr := miniredis.RunT(t)
    rdb := redis.NewClient(&redis.Options{Addr: mr.Addr()})

    svc := &UserService{rdb: rdb, origin: fakeOrigin{}}
    body, err := svc.GetUser(context.Background(), "42")
    require.NoError(t, err)
    assert.True(t, mr.Exists("bff:user:42"))
}
```

Mock `OriginClient` with an interface — same pattern as Chapter 15 AWS clients.

---

## Key Takeaways

1. **One `*redis.Client` per process** — pooled connections; inject into services.
2. **Always pass `context.Context`** — timeouts prevent Redis slowness from stalling the BFF.
3. **Treat `redis.Nil` as cache miss** — distinct from infrastructure errors.
4. **Pipelines reduce RTT** — batch GETs during aggregation endpoints.
5. **Cache-aside: origin is source of truth** — SET failures are logged, not fatal.
6. **Match TTL to freshness needs** — BFF caches are usually short-lived.

---

**Practice**: [Redis Get/Set →](../exercises/module-11-data-messaging/61-redis-get-set/) | [Cache-Aside →](../exercises/module-11-data-messaging/62-redis-cache-aside/) | [Module 07 — Timeouts →](../exercises/module-07-resilience/46-timeout-patterns/)
