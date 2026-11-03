# Exercise 42: fiber-error-handling

## Problem

Centralize HTTP error responses with a custom Fiber `ErrorHandler`.

Requirements:

1. Create the app with **`fiber.New(fiber.Config{ErrorHandler: ...})`**.
2. The error handler returns JSON **`{"error": "<message>"}`**. Use the status code from `*fiber.Error` when the error is a Fiber error; otherwise use **500**.
3. Routes:
   - `GET /boom` → return `fiber.NewError(fiber.StatusBadRequest, "bad request")`
   - `GET /panic-ish` → return `errors.New("unexpected")` (handled as 500)

## How to Test

```bash
go test -v
```
