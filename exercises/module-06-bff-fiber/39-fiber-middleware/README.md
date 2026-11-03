# Exercise 39: fiber-middleware

## Problem

Implement custom Fiber middleware that ensures every response includes an `X-Request-ID` header.

Requirements:

1. **`RequestIDMiddleware(generate RequestIDGenerator) fiber.Handler`** — If the incoming request has no `X-Request-ID` header, generate one with `generate` (when `generate` is nil, use a simple default generator). Set `X-Request-ID` on the response. If the client already sent `X-Request-ID`, echo that value on the response. Call `c.Next()`.
2. **`SetupApp()`** — Register the middleware and a `GET /ping` route that returns `"pong"`.

## How to Test

```bash
go test -v
```
