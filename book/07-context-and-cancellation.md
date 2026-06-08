# Chapter 07: Context & Cancellation

> Context is how Go propagates deadlines, cancellation signals, and request-scoped values.

---

## Why Context?

In Java, you might use:
- `Thread.interrupt()` for cancellation
- `HttpServletRequest` for request-scoped data
- Timeout parameters on methods

Go consolidates all of this into the `context` package.

```java
// Java: Scattered cancellation mechanisms
public CompletableFuture<String> fetch(String url) {
    return CompletableFuture.supplyAsync(() -> {
        // How do you cancel this? Check Thread.interrupted()?
        // How do you propagate a deadline?
    });
}
```

```go
// Go: Context propagates through the call stack
func fetch(ctx context.Context, url string) (string, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return "", err
    }
    resp, err := http.DefaultClient.Do(req)
    // Automatically cancelled if ctx is cancelled/timed out
}
```

---

## Context Basics

### Creating Contexts

```go
// Background context — root of all context trees
ctx := context.Background()

// TODO context — when unsure which context to use
ctx := context.TODO()
```

### With Cancellation

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()  // Always call cancel to free resources!

go func() {
    <-ctx.Done()  // Blocks until cancelled
    fmt.Println("cancelled:", ctx.Err())  // "context canceled"
}()

// Cancel after some condition
time.Sleep(2 * time.Second)
cancel()
```

### With Timeout

```go
// Auto-cancel after 5 seconds
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

result, err := doSlowOperation(ctx)
if err != nil {
    if ctx.Err() == context.DeadlineExceeded {
        // Timeout
    }
}
```

### With Deadline

```go
// Cancel at specific time
deadline := time.Now().Add(10 * time.Second)
ctx, cancel := context.WithDeadline(context.Background(), deadline)
defer cancel()
```

---

## Propagating Context

### The Rule: Context flows down the call stack

```go
func Handler(ctx context.Context) {
    // Pass ctx to all children
    user, err := getUser(ctx, userID)
    if err != nil {
        return
    }
    orders, err := getOrders(ctx, user.ID)
    // ...
}

func getUser(ctx context.Context, id int) (*User, error) {
    // ctx carries deadline/cancellation from handler
    req, _ := http.NewRequestWithContext(ctx, "GET", "/users/"+strconv.Itoa(id), nil)
    resp, err := http.DefaultClient.Do(req)
    // If handler's ctx times out, this request is automatically cancelled
}

func getOrders(ctx context.Context, userID int) ([]Order, error) {
    // Same ctx — same deadline/cancellation
}
```

### Database Queries with Context

```go
func getUser(ctx context.Context, db *sql.DB, id int) (*User, error) {
    row := db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = ?", id)
    var user User
    err := row.Scan(&user.ID, &user.Name)
    return &user, err
}
```

---

## Context Values — Request-Scoped Data

### Basic Usage

```go
type key string
const requestIDKey key = "request_id"

// Set value
ctx := context.WithValue(context.Background(), requestIDKey, "abc-123")

// Get value
if id, ok := ctx.Value(requestIDKey).(string); ok {
    fmt.Println("Request ID:", id)
}
```

### Best Practice: Define Custom Types

```go
// Define a custom key type to avoid collisions
type contextKey string

const (
    RequestIDKey contextKey = "request_id"
    UserIDKey    contextKey = "user_id"
)

// Helper functions
func WithRequestID(ctx context.Context, id string) context.Context {
    return context.WithValue(ctx, RequestIDKey, id)
}

func GetRequestID(ctx context.Context) string {
    if id, ok := ctx.Value(RequestIDKey).(string); ok {
        return id
    }
    return ""
}
```

> ⚠️ **Rule**: Don't use context values for optional parameters or business logic. Use them only for cross-cutting concerns (request IDs, tracing IDs, auth tokens).

---

## Context in HTTP Servers

### Fiber Context (BFF Pattern)

```go
func (h *Handler) GetUser(c *fiber.Ctx) error {
    // Fiber has its ownCtx, but bridge to stdlib context
    ctx := c.Context()
    
    // Set timeout
    ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
    defer cancel()
    
    user, err := h.service.GetUser(ctx, c.Params("id"))
    if err != nil {
        return err
    }
    return c.JSON(user)
}
```

### Middleware Setting Context

```go
func requestIDMiddleware(c *fiber.Ctx) error {
    id := uuid.New().String()
    c.Locals("request_id", id)  // Fiber's equivalent of context value
    return c.Next()
}
```

---

## Common Patterns

### Pattern: Goroutine with Context Cancellation

```go
func process(ctx context.Context, data <-chan Data) error {
    for {
        select {
        case <-ctx.Done():
            return ctx.Err()  // Clean shutdown
        case d := <-data:
            if err := handle(ctx, d); err != nil {
                return err
            }
        }
    }
}
```

### Pattern: Context Timeout for External Calls

```go
func (s *Service) CallDownstream(ctx context.Context, req Request) (Response, error) {
    // Always set a timeout for external calls
    ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()
    
    // If parent ctx has shorter deadline, that takes precedence
    body, _ := json.Marshal(req)
    httpReq, _ := http.NewRequestWithContext(ctx, "POST", s.url, bytes.NewReader(body))
    
    resp, err := s.client.Do(httpReq)
    if err != nil {
        return Response{}, fmt.Errorf("downstream call: %w", err)
    }
    defer resp.Body.Close()
    
    // ...
}
```

---

## Key Takeaways

1. **Context flows down** — Always pass as first parameter
2. **Always call `cancel()`** — Use `defer cancel()` immediately after creation
3. **Use `WithTimeout` for external calls** — Never let requests hang indefinitely
4. **Don't store business data in context** — Only cross-cutting concerns
5. **Check `ctx.Done()` in goroutines** — For clean cancellation
6. **`context.Background()` is the root** — Never use `TODO` in production

---

**Practice**: [Module 04 — Concurrency Exercises →](../exercises/module-04-concurrency/)

Next: [Chapter 08 — Design Patterns →](08-design-patterns.md)