# Exercises: Java to Go — Hands-On Practice

> 66 exercises organized in 11 modules, from Go basics to data & messaging patterns.

## How Exercises Work

Each exercise folder contains:
- **`README.md`** — Problem description and learning objectives
- **`solution.go`** — Broken/stub code for you to fix
- **`solution_test.go`** — Tests that validate your fix
- **`_solution/`** — Working solution (cheating! 😄)

### Workflow
```bash
cd exercises/module-01-basics/01-hello-go
# Read the README
cat README.md
# Read the broken code
cat solution.go
# Fix the code, then run tests
go test -v
# If tests pass, you're done!
```

---

## Module 01: Go Basics (Exercises 01-08)
| # | Exercise | Topic |
|---|----------|-------|
| 01 | [hello-go](module-01-basics/01-hello-go/) | First Go program |
| 02 | [variables-and-types](module-01-basics/02-variables-and-types/) | Variables, zero values, structs |
| 03 | [constants-and-enums](module-01-basics/03-constants-and-enums/) | Constants, custom types |
| 04 | [functions-and-multiple-returns](module-01-basics/04-functions-and-multiple-returns/) | Multiple returns, errors |
| 05 | [pointers-and-references](module-01-basics/05-pointers-and-references/) | Pointers, swap |
| 06 | [structs-and-composition](module-01-basics/06-structs-and-composition/) | Struct methods |
| 07 | [methods-and-receivers](module-01-basics/07-methods-and-receivers/) | Value vs pointer receivers |
| 08 | [interfaces](module-01-basics/08-interfaces/) | Implicit interfaces |

## Module 02: Collections (Exercises 09-14)
| # | Exercise | Topic |
|---|----------|-------|
| 09 | [slice-operations](module-02-collections/09-slice-operations/) | Filtering slices |
| 10 | [slice-append-and-capacity](module-02-collections/10-slice-append-and-capacity/) | Pre-allocation |
| 11 | [map-crud](module-02-collections/11-map-crud/) | Key-value store |
| 12 | [map-iteration-order](module-02-collections/12-map-iteration-order/) | Sorted keys |
| 13 | [embedding-in-structs](module-02-collections/13-embedding-in-structs/) | Struct embedding |
| 14 | [type-conversion](module-02-collections/14-type-conversion/) | Type conversion |

## Module 03: Error Handling (Exercises 15-20)
| # | Exercise | Topic |
|---|----------|-------|
| 15 | [basic-errors](module-03-error-handling/15-basic-errors/) | errors.New |
| 16 | [custom-error-types](module-03-error-handling/16-custom-error-types/) | Custom error types |
| 17 | [error-wrapping-and-unwrapping](module-03-error-handling/17-error-wrapping-and-unwrapping/) | fmt.Errorf %w |
| 18 | [errors-is-and-as](module-03-error-handling/18-errors-is-and-as/) | Error chain inspection |
| 19 | [panic-and-recover](module-03-error-handling/19-panic-and-recover/) | Panic recovery |
| 20 | [sentinel-errors](module-03-error-handling/20-sentinel-errors/) | Sentinel pattern |

## Module 04: Concurrency (Exercises 21-30)
| # | Exercise | Topic |
|---|----------|-------|
| 21 | [goroutine-basics](module-04-concurrency/21-goroutine-basics/) | Parallel sum |
| 22 | [channels-unbuffered](module-04-concurrency/22-channels-unbuffered/) | Producer-consumer |
| 23 | [channels-buffered](module-04-concurrency/23-channels-buffered/) | Semaphore pattern |
| 24 | [select-statement](module-04-concurrency/24-select-statement/) | Channel multiplexing |
| 25 | [fan-out-fan-in](module-04-concurrency/25-fan-out-fan-in/) | Fan-out/fan-in |
| 26 | [context-timeout](module-04-concurrency/26-context-timeout/) | Context timeout |
| 27 | [context-cancellation](module-04-concurrency/27-context-cancellation/) | Cancellation propagation |
| 28 | [context-values](module-04-concurrency/28-context-values/) | Request-scoped values |
| 29 | [waitgroup](module-04-concurrency/29-waitgroup/) | Goroutine synchronization |
| 30 | [mutex-and-data-races](module-04-concurrency/30-mutex-and-data-races/) | Thread safety |

## Module 05: Testing (Exercises 31-36)
| # | Exercise | Topic |
|---|----------|-------|
| 31 | [unit-test-basics](module-05-testing/31-unit-test-basics/) | Basic testing |
| 32 | [table-driven-tests](module-05-testing/32-table-driven-tests/) | Table-driven pattern |
| 33 | [sub-tests](module-05-testing/33-sub-tests/) | t.Run organization |
| 34 | [mocking-interfaces](module-05-testing/34-mocking-interfaces/) | Interface mocking |
| 35 | [benchmarking](module-05-testing/35-benchmarking/) | strings.Builder optimization |
| 36 | [http-handler-testing](module-05-testing/36-http-handler-testing/) | httptest package |

## Module 06: BFF with Fiber (Exercises 37-42)
| # | Exercise | Topic |
|---|----------|-------|
| 37 | [fiber-hello-world](module-06-bff-fiber/37-fiber-hello-world/) | First Fiber app |
| 38 | [fiber-routing-and-params](module-06-bff-fiber/38-fiber-routing-and-params/) | Path/query params |
| 39 | [fiber-middleware](module-06-bff-fiber/39-fiber-middleware/) | X-Request-ID middleware |
| 40 | [fiber-json-binding](module-06-bff-fiber/40-fiber-json-binding/) | BodyParser + 201/400 |
| 41 | [fiber-group-routes](module-06-bff-fiber/41-fiber-group-routes/) | `/api/v1` route groups |
| 42 | [fiber-error-handling](module-06-bff-fiber/42-fiber-error-handling/) | JSON ErrorHandler |

## Module 07: Resilience Patterns (Exercises 43-48)
| # | Exercise | Topic |
|---|----------|-------|
| 43 | [circuit-breaker](module-07-resilience/43-circuit-breaker/) | Circuit breaker |
| 44 | [retry-with-backoff](module-07-resilience/44-retry-with-backoff/) | Exponential backoff |
| 45 | [rate-limiting](module-07-resilience/45-rate-limiting/) | Token bucket limiter |
| 46 | [timeout-patterns](module-07-resilience/46-timeout-patterns/) | Context timeout wrapper |
| 47 | [graceful-shutdown](module-07-resilience/47-graceful-shutdown/) | Cancel + WaitGroup shutdown |
| 48 | [health-check-endpoint](module-07-resilience/48-health-check-endpoint/) | `/health` + `/ready` |

## Module 08: Observability (Exercises 49-53)
| # | Exercise | Topic |
|---|----------|-------|
| 49 | [structured-logging](module-08-observability/49-structured-logging/) | slog JSON logging |
| 50 | [request-tracing](module-08-observability/50-request-tracing/) | X-Request-ID middleware |
| 51 | [middleware-logging](module-08-observability/51-middleware-logging/) | Access logs with slog |
| 52 | [distributed-tracing](module-08-observability/52-distributed-tracing/) | Tracer interface (Datadog-style) |
| 53 | [metrics-collection](module-08-observability/53-metrics-collection/) | Labelled HTTP counters |

## Module 09: Configuration with Viper (Exercises 54-56)
| # | Exercise | Topic |
|---|----------|-------|
| 54 | [environment-config](module-09-viper-config/54-environment-config/) | Env var config |
| 55 | [hot-reload-configmap](module-09-viper-config/55-hot-reload-configmap/) | Viper YAML/JSON hot reload |
| 56 | [feature-flags](module-09-viper-config/56-feature-flags/) | Dynamic feature flags |

## Module 10: Advanced BFF Patterns (Exercises 57-60)
| # | Exercise | Topic |
|---|----------|-------|
| 57 | [parallel-http-requests](module-10-advanced/57-parallel-http-requests/) | Concurrent FetchAll + cancel |
| 58 | [response-aggregation](module-10-advanced/58-response-aggregation/) | Parallel dashboard aggregate |
| 59 | [pooling-and-reuse](module-10-advanced/59-pooling-and-reuse/) | sync.Pool buffer reuse |
| 60 | [full-bff-integration](module-10-advanced/60-full-bff-integration/) | Fiber BFF + cache-aside |

## Module 11: Data & Messaging (Exercises 61-66)
| # | Exercise | Topic |
|---|----------|-------|
| 61 | [redis-get-set](module-11-data-messaging/61-redis-get-set/) | Cache Client / Get-Set |
| 62 | [redis-cache-aside](module-11-data-messaging/62-redis-cache-aside/) | GetOrLoad cache-aside |
| 63 | [dynamodb-put-get](module-11-data-messaging/63-dynamodb-put-get/) | DynamoDB Put/Get (in-memory) |
| 64 | [dynamodb-query](module-11-data-messaging/64-dynamodb-query/) | Query by partition key |
| 65 | [sns-publish](module-11-data-messaging/65-sns-publish/) | SNS Publish interface |
| 66 | [event-outbox-pattern](module-11-data-messaging/66-event-outbox-pattern/) | Transactional outbox |

---

**Tip**: Start from Module 01 and work through sequentially. Each module builds on concepts from previous ones.

← [Back to Study Book](../book/00-introduction.md)
