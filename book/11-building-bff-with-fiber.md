# Chapter 11: Building a BFF with Fiber

> Fiber is an Express.js-inspired web framework for Go. Fast, simple, and perfect for BFF patterns.

---

## Why Fiber?

| Feature | Java Spring Boot | Go Fiber |
|---------|-----------------|----------|
| Startup | 5-15 seconds | ~50ms |
| Memory | 200-500 MB | ~10 MB |
| Throughput | ~50K req/s | ~200K+ req/s |
| Routing | Annotation-based | Explicit |
| Middleware | Filter/Interceptor chains | Simple function chains |

---

## Setup

```bash
mkdir bff-service && cd bff-service
go mod init github.com/yourname/bff-service
go get github.com/gofiber/fiber/v2
```

## Hello World

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()
    
    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, BFF!")
    })
    
    app.Listen(":3000")
}
```

---

## Routing

```go
app.Get("/users/:id", getUser)          // Path parameter
app.Post("/users", createUser)          // POST
app.Put("/users/:id", updateUser)       // PUT
app.Delete("/users/:id", deleteUser)    // DELETE
app.Get("/search", search)              // Query params

func getUser(c *fiber.Ctx) error {
    id := c.Params("id")               // Path param
    return c.JSON(fiber.Map{"id": id})
}

func search(c *fiber.Ctx) error {
    q := c.Query("q")                  // Query param
    page := c.QueryInt("page", 1)      // With default
    return c.JSON(fiber.Map{"q": q, "page": page})
}
```

---

## JSON Binding

```go
type CreateUserRequest struct {
    Name  string `json:"name"`
    Email string `json:"email"`
    Age   int    `json:"age"`
}

func createUser(c *fiber.Ctx) error {
    req := new(CreateUserRequest)
    if err := c.BodyParser(req); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "invalid body"})
    }
    // Use req.Name, req.Email, req.Age
    return c.Status(201).JSON(req)
}
```

---

## Middleware

```go
// Custom middleware
func LoggerMiddleware(c *fiber.Ctx) error {
    start := time.Now()
    err := c.Next()  // Call next handler
    duration := time.Since(start)
    log.Info("request", "method", c.Method(), "path", c.Path(), "duration", duration)
    return err
}

// Rate limiting middleware
func RateLimitMiddleware(maxRequests int, window time.Duration) fiber.Handler {
    // Implementation using token bucket
}

// Apply middleware
app.Use(LoggerMiddleware)
app.Use(RateLimitMiddleware(100, time.Minute))

// Route-specific middleware
app.Get("/admin", AuthMiddleware, adminHandler)
```

---

## Route Groups

```go
api := app.Group("/api/v1")

users := api.Group("/users")
users.Get("/", listUsers)
users.Get("/:id", getUser)
users.Post("/", createUser)

orders := api.Group("/orders")
orders.Get("/", listOrders)
orders.Post("/", createOrder)
```

---

## Error Handling

```go
// Centralized error handler
app.Use(func(c *fiber.Ctx) error {
    err := c.Next()
    if err != nil {
        code := fiber.StatusInternalServerError
        if e, ok := err.(*fiber.Error); ok {
            code = e.Code
        }
        return c.Status(code).JSON(fiber.Map{
            "error": err.Error(),
        })
    }
    return nil
})
```

---

## BFF Pattern: Aggregating Downstream Calls

```go
type BFFHandler struct {
    userService    UserService
    orderService   OrderService
    paymentService PaymentService
}

func (h *BFFHandler) GetUserDashboard(c *fiber.Ctx) error {
    userID := c.Params("id")
    ctx, cancel := context.WithTimeout(c.Context(), 3*time.Second)
    defer cancel()
    
    // Parallel calls using errgroup
    g, ctx := errgroup.WithContext(ctx)
    
    var user *User
    var orders []Order
    var payments []Payment
    
    g.Go(func() error {
        var err error
        user, err = h.userService.GetUser(ctx, userID)
        return err
    })
    g.Go(func() error {
        var err error
        orders, err = h.orderService.GetOrders(ctx, userID)
        return err
    })
    g.Go(func() error {
        var err error
        payments, err = h.paymentService.GetPayments(ctx, userID)
        return err
    })
    
    if err := g.Wait(); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    
    return c.JSON(fiber.Map{
        "user":     user,
        "orders":   orders,
        "payments": payments,
    })
}
```

---

## Graceful Shutdown

```go
func main() {
    app := fiber.New()
    // ... setup routes
    
    // Graceful shutdown
    go func() {
        sigChan := make(chan os.Signal, 1)
        signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
        <-sigChan
        log.Info("Shutting down...")
        app.ShutdownWithTimeout(10 * time.Second)
    }()
    
    if err := app.Listen(":3000"); err != nil {
        log.Error("server stopped", "error", err)
    }
}
```

---

## Key Takeaways

1. **Fiber is fast** — Built on fasthttp, not net/http
2. **Middleware chains** — Like Express.js, familiar for Node devs too
3. **Route groups** — Organize APIs cleanly
4. **Parallel aggregation** — BFF pattern with errgroup
5. **Always graceful shutdown** — Don't drop in-flight requests
6. **Context with timeouts** — Every downstream call needs one

---

**Practice**: [Module 06 — BFF/Fiber Exercises →](../exercises/module-06-bff-fiber/) | [Module 07 — Resilience →](../exercises/module-07-resilience/) | [Module 10 — Advanced →](../exercises/module-10-advanced/)