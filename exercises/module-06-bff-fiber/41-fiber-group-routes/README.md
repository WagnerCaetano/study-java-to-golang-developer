# Exercise 41: fiber-group-routes

## Problem

Organize routes with Fiber groups and subgroups.

Requirements for **`SetupApp()`**:

1. Create a group prefix **`/api/v1`**.
2. Under it, a subgroup **`/users`** with:
   - `GET /` → JSON `{"users":[]}`
   - `GET /:id` → JSON `{"id":"<param>"}` (id from the path)

## How to Test

```bash
go test -v
```
