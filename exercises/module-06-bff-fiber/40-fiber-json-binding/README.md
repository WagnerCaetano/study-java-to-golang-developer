# Exercise 40: fiber-json-binding

## Problem

Parse JSON request bodies with Fiber and return structured JSON responses.

Requirements:

1. Define **`CreateUserRequest`** with `Name string \`json:"name"\`` and `Email string \`json:"email"\``.
2. **`SetupApp()`** — `POST /users` uses `BodyParser` on the request. Valid JSON → **201** with the parsed struct echoed as JSON. Invalid JSON → **400** with `{"error":"invalid body"}`.

## How to Test

```bash
go test -v
```
