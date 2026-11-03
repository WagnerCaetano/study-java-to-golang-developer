# Exercise 52: Distributed Tracing (Datadog-style)

## Problem

**Distributed tracing** ties together spans across services. In production you might use [Datadog dd-trace-go v2](https://github.com/DataDog/dd-trace-go) (open source; the Datadog backend is optional). This exercise uses a **small `Tracer` interface** so tests run without an agent.

Package **`ddtraceex`**:

1. `Tracer` with `StartSpan(name string) (finish func())`.
2. `RecordingTracer` — appends each started span name to `Spans` (finish is a no-op).
3. `NoopTracer` — no-op spans.
4. `SetDefaultTracer(t Tracer)` and `StartSpan(name string) (finish func())` delegating to the default (initially `NoopTracer`).

Implement `RecordingTracer.StartSpan` so the name is recorded when the span **starts**.

## Learning objectives

- Wrap vendor SDKs behind interfaces for testability.
- Span lifecycle (`start` / `finish`).
- How BFF handlers create child spans around downstream calls.

**Real library:** `github.com/DataDog/dd-trace-go/v2/ddtrace/tracer` — wire `DatadogTracer` in your app when you have an agent or serverless integration.

## How to test

```bash
go test -v
```
