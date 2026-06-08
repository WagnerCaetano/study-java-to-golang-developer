# Java → Go Developer Study Book

> A comprehensive study guide for Java developers migrating to Go, focused on **high-traffic backend development**.

---

## 📖 Study Book — Overview Chapters

The book provides a guided tour of Go from a Java developer's perspective, with side-by-side comparisons, memory model insights, and production-ready patterns.

| # | Chapter | Description |
|---|---------|-------------|
| 00 | [Introduction](book/00-introduction.md) | Why Go? Java vs Go philosophy, tooling setup |
| 01 | [Language Fundamentals](book/01-language-fundamentals.md) | Variables, types, constants, functions |
| 02 | [Structs & Methods](book/02-structs-and-methods.md) | Go structs vs Java classes, method receivers, pointers |
| 03 | [Interfaces & Polymorphism](book/03-interfaces-and-polymorphism.md) | Implicit interfaces, duck typing vs explicit implements |
| 04 | [Arrays, Slices & Maps](book/04-arrays-slices-maps.md) | Collection types, memory layout, capacity vs length |
| 05 | [Error Handling](book/05-error-handling.md) | No exceptions — Go's error philosophy |
| 06 | [Concurrency Model](book/06-concurrency-model.md) | Goroutines, channels, scheduler internals |
| 07 | [Context & Cancellation](book/07-context-and-cancellation.md) | Context propagation patterns |
| 08 | [Design Patterns](book/08-design-patterns.md) | Decoupling, mocking, type grouping, conversion |
| 09 | [Testing & Benchmarking](book/09-testing-and-benchmarking.md) | Go testing philosophy, table-driven tests |
| 10 | [Profiling & Optimization](book/10-profiling-and-optimization.md) | Memory tracing, stack traces, escape analysis |
| 11 | [Building a BFF with Fiber](book/11-building-bff-with-fiber.md) | Fiber framework, middleware, routing patterns |

---

## 🧪 Hands-On Exercises — 60 Scenarios

Exercises follow an **Exercism-style** format: read the problem, fix the broken code, run tests to validate.

### Module 01 — Go Basics (Exercises 01–08)
| # | Exercise | Topic |
|---|----------|-------|
| 01 | [hello-go](exercises/module-01-basics/01-hello-go/README.md) | Your first Go program |
| 02 | [variables-and-types](exercises/module-01-basics/02-variables-and-types/README.md) | Variable declaration, zero values |
| 03 | [constants-and-enums](exercises/module-01-basics/03-constants-and-enums/README.md) | Constants and iota enums |
| 04 | [functions-and-multiple-returns](exercises/module-01-basics/04-functions-and-multiple-returns/README.md) | Multiple return values |
| 05 | [pointers-and-references](exercises/module-01-basics/05-pointers-and-references/README.md) | Pointers vs Java references |
| 06 | [structs-and-composition](exercises/module-01-basics/06-structs-and-composition/README.md) | Struct composition |
| 07 | [methods-and-receivers](exercises/module-01-basics/07-methods-and-receivers/README.md) | Value vs pointer receivers |
| 08 | [interfaces](exercises/module-01-basics/08-interfaces/README.md) | Implicit interface implementation |

### Module 02 — Collections (Exercises 09–14)
| # | Exercise | Topic |
|---|----------|-------|
| 09 | [slice-operations](exercises/module-02-collections/09-slice-operations/README.md) | Slice basics and indexing |
| 10 | [slice-append-and-capacity](exercises/module-02-collections/10-slice-append-and-capacity/README.md) | Append, capacity, copy |
| 11 | [map-crud](exercises/module-02-collections/11-map-crud/README.md) | Map create, read, update, delete |
| 12 | [map-iteration-order](exercises/module-02-collections/12-map-iteration-order/README.md) | Map iteration randomness |
| 13 | [embedding-in-structs](exercises/module-02-collections/13-embedding-in-structs/README.md) | Struct embedding |
| 14 | [type-conversion](exercises/module-02-collections/14-type-conversion/README.md) | Type conversion and assertion |

### Module 03 — Error Handling (Exercises 15–20)
| # | Exercise | Topic |
|---|----------|-------|
| 15 | [basic-errors](exercises/module-03-error-handling/15-basic-errors/README.md) | Error creation and checking |
| 16 | [custom-error-types](exercises/module-03-error-handling/16-custom-error-types/README.md) | Custom error structs |
| 17 | [error-wrapping-and-unwrapping](exercises/module-03-error-handling/17-error-wrapping-and-unwrapping/README.md) | fmt.Errorf with %w |
| 18 | [errors-is-and-as](exercises/module-03-error-handling/18-errors-is-and-as/README.md) | errors.Is and errors.As |
| 19 | [panic-and-recover](exercises/module-03-error-handling/19-panic-and-recover/README.md) | Panic/recover patterns |
| 20 | [sentinel-errors](exercises/module-03-error-handling/20-sentinel-errors/README.md) | Sentinel error pattern |

### Module 04 — Concurrency (Exercises 21–30)
| # | Exercise | Topic |
|---|----------|-------|
| 21 | [goroutine-basics](exercises/module-04-concurrency/21-goroutine-basics/README.md) | Launching goroutines |
| 22 | [channels-unbuffered](exercises/module-04-concurrency/22-channels-unbuffered/README.md) | Unbuffered channel signaling |
| 23 | [channels-buffered](exercises/module-04-concurrency/23-channels-buffered/README.md) | Buffered channels |
| 24 | [select-statement](exercises/module-04-concurrency/24-select-statement/README.md) | Multiplexing with select |
| 25 | [fan-out-fan-in](exercises/module-04-concurrency/25-fan-out-fan-in/README.md) | Fan-out/fan-in concurrency |
| 26 | [context-timeout](exercises/module-04-concurrency/26-context-timeout/README.md) | Context with timeout |
| 27 | [context-cancellation](exercises/module-04-concurrency/27-context-cancellation/README.md) | Context cancellation propagation |
| 28 | [context-values](exercises/module-04-concurrency/28-context-values/README.md) | Context value propagation |
| 29 | [waitgroup](exercises/module-04-concurrency/29-waitgroup/README.md) | Synchronizing goroutines |
| 30 | [mutex-and-data-races](exercises/module-04-concurrency/30-mutex-and-data-races/README.md) | Mutex, race detection |

### Module 05 — Testing (Exercises 31–36)
| # | Exercise | Topic |
|---|----------|-------|
| 31 | [unit-test-basics](exercises/module-05-testing/31-unit-test-basics/README.md) | Basic test functions |
| 32 | [table-driven-tests](exercises/module-05-testing/32-table-driven-tests/README.md) | Table-driven test pattern |
| 33 | [sub-tests](exercises/module-05-testing/33-sub-tests/README.md) | Sub-tests with t.Run |
| 34 | [mocking-interfaces](exercises/module-05-testing/34-mocking-interfaces/README.md) | Mocking with interfaces |
| 35 | [benchmarking](exercises/module-05-testing/35-benchmarking/README.md) | Writing benchmarks |
| 36 | [http-handler-testing](exercises/module-05-testing/36-http-handler-testing/README.md) | Testing HTTP handlers |

### Module 06 — BFF with Fiber (Exercises 37–42)
| # | Exercise | Topic |
|---|----------|-------|
| 37 | [fiber-hello-world](exercises/module-06-bff-fiber/37-fiber-hello-world/README.md) | First Fiber app |
| 38 | [fiber-routing-and-params](exercises/module-06-bff-fiber/38-fiber-routing-and-params/README.md) | Routes and parameters |
| 39 | [fiber-middleware](exercises/module-06-bff-fiber/39-fiber-middleware/README.md) | Custom middleware |
| 40 | [fiber-json-binding](exercises/module-06-bff-fiber/40-fiber-json-binding/README.md) | JSON request/response |
| 41 | [fiber-group-routes](exercises/module-06-bff-fiber/41-fiber-group-routes/README.md) | Route grouping |
| 42 | [fiber-error-handling](exercises/module-06-bff-fiber/42-fiber-error-handling/README.md) | Centralized error handling |

### Module 07 — Resilience Patterns (Exercises 43–48)
| # | Exercise | Topic |
|---|----------|-------|
| 43 | [circuit-breaker](exercises/module-07-resilience/43-circuit-breaker/README.md) | Circuit breaker implementation |
| 44 | [retry-with-backoff](exercises/module-07-resilience/44-retry-with-backoff/README.md) | Exponential backoff retry |
| 45 | [rate-limiting](exercises/module-07-resilience/45-rate-limiting/README.md) | Token bucket rate limiting |
| 46 | [timeout-patterns](exercises/module-07-resilience/46-timeout-patterns/README.md) | Timeout patterns for HTTP calls |
| 47 | [graceful-shutdown](exercises/module-07-resilience/47-graceful-shutdown/README.md) | Graceful server shutdown |
| 48 | [health-check-endpoint](exercises/module-07-resilience/48-health-check-endpoint/README.md) | Health check endpoints |

### Module 08 — Observability (Exercises 49–53)
| # | Exercise | Topic |
|---|----------|-------|
| 49 | [structured-logging](exercises/module-08-observability/49-structured-logging/README.md) | Structured logging with slog |
| 50 | [request-tracing](exercises/module-08-observability/50-request-tracing/README.md) | Request ID middleware |
| 51 | [middleware-logging](exercises/module-08-observability/51-middleware-logging/README.md) | Request/response logging middleware |
| 52 | [distributed-tracing](exercises/module-08-observability/52-distributed-tracing/README.md) | OpenTelemetry tracing |
| 53 | [metrics-collection](exercises/module-08-observability/53-metrics-collection/README.md) | Prometheus metrics |

### Module 09 — Kubernetes Dynamic Configuration (Exercises 54–56)
| # | Exercise | Topic |
|---|----------|-------|
| 54 | [environment-config](exercises/module-09-kubernetes-config/54-environment-config/README.md) | Environment-based config |
| 55 | [hot-reload-configmap](exercises/module-09-kubernetes-config/55-hot-reload-configmap/README.md) | ConfigMap hot reload |
| 56 | [feature-flags](exercises/module-09-kubernetes-config/56-feature-flags/README.md) | Dynamic feature flags |

### Module 10 — Advanced BFF Patterns (Exercises 57–60)
| # | Exercise | Topic |
|---|----------|-------|
| 57 | [parallel-http-requests](exercises/module-10-advanced/57-parallel-http-requests/README.md) | Concurrent API calls |
| 58 | [response-aggregation](exercises/module-10-advanced/58-response-aggregation/README.md) | Aggregating parallel responses |
| 59 | [pooling-and-reuse](exercises/module-10-advanced/59-pooling-and-reuse/README.md) | Connection pooling, sync.Pool |
| 60 | [full-bff-integration](exercises/module-10-advanced/60-full-bff-integration/README.md) | Complete BFF with all patterns |

---

## 🚀 How to Use

1. **Read the book chapters** in order (00 → 11) to build understanding
2. **Work through exercises** module by module
3. For each exercise:
   - Read the `README.md` for the problem description
   - Fix the code in the stub file
   - Run `go test` to validate your solution
   - Check `_solution/` if stuck

## 🛠 Prerequisites

- Go 1.21+ installed
- Basic Java knowledge
- Familiarity with backend development concepts

```bash
# Verify Go installation
go version

# Run any exercise test
cd exercises/module-01-basics/01-hello-go
go test -v
```

---

*Based on [Ultimate Go by Ardan Labs](https://github.com/hoanhan101/ultimate-go)*