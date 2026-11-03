# Exercise 49: Structured Logging

## Problem

Production services emit **structured logs** (JSON lines) so log aggregators can index fields. Go’s standard library `log/slog` provides levelled logging with typed attributes.

Implement:

1. `NewJSONLogger(w io.Writer) *slog.Logger` — logger whose handler writes **JSON** (not text).
2. `LogInfo(logger, msg, attrs...)` — log at Info with the given `slog.Attr` values.
3. `LogError(logger, msg, err, attrs...)` — log at Error, including the error as an attribute named `error`, plus any extra attrs.

Package name: **`logging`**.

## Learning objectives

- Configure `slog.NewJSONHandler` and log levels.
- Use `LogAttrs` (or equivalent) to attach structured fields.
- Parse JSON log output in tests (mirrors how Loki/CloudWatch ingest logs).

## How to test

```bash
go test -v
```
