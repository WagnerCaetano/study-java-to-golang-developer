# Exercise 50: Request Tracing (Request ID)

## Problem

BFFs and APIs propagate a **correlation ID** so logs and traces for one HTTP call line up. Implement Fiber middleware that ensures every request has an `X-Request-ID` header.

Implement in package **`reqtrace`**:

1. `GenerateRequestID() string` — cryptographically random **16-byte** value encoded as **hex** (32 characters).
2. `RequestIDMiddleware() fiber.Handler` — if the client did not send `X-Request-ID`, generate one; store it on the context (Fiber locals) and set the response header.
3. `GetRequestID(c *fiber.Ctx) string` — read the ID from locals (empty if missing).

Constant: `HeaderRequestID = "X-Request-ID"`.

## Learning objectives

- Fiber middleware and `c.Locals`.
- `crypto/rand` for request IDs (avoid predictable IDs in production).
- Propagate correlation headers across services.

## How to test

```bash
go test -v
```
