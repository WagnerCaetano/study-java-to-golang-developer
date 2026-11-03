# Exercise 51: Middleware Logging

## Problem

Every HTTP request should leave an audit-friendly log line: method, path, status, and duration. Implement Fiber middleware that logs those fields with **`slog`**.

Package **`mwlog`**:

- `LoggingMiddleware(logger *slog.Logger) fiber.Handler` — call `c.Next()`, then log one line with keys `method`, `path`, `status`, and `duration_ms` (integer milliseconds).

Use the provided logger; do not replace the global default logger.

## Learning objectives

- Measure handler latency in middleware.
- Structured access logs (same shape as Java Spring `CommonsRequestLoggingFilter` or similar).
- Compose middleware with routes under test.

## How to test

```bash
go test -v
```
