# Java → Go Developer Study Book

> A comprehensive study guide for Java developers migrating to Go, focused on **high-traffic BFF / backend development** with a production-ready stack.

**Target stack:** Go **1.25+** · Fiber **v2** · Viper (YAML+JSON + fsnotify hot reload) · Datadog **dd-trace-go v2** (OSS) · Prometheus (**fiberprometheus/v2**) · Redis (**go-redis/v9**) · DynamoDB & SNS (**aws-sdk-go-v2**) · Feature flags via Viper

---

## Study Book — Chapters

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
| 11 | [Building a BFF with Fiber](book/11-building-bff-with-fiber.md) | Fiber v2 framework, middleware, routing patterns |
| 12 | [Configuration with Viper](book/12-configuration-with-viper.md) | YAML + JSON, env overrides, fsnotify hot reload |
| 13 | [Observability: Datadog & Prometheus](book/13-observability-datadog-prometheus.md) | dd-trace-go v2 + fiberprometheus/v2 |
| 14 | [Redis with go-redis](book/14-redis-with-go-redis.md) | Caching, TTL, pipelines, cache-aside |
| 15 | [AWS DynamoDB & SNS](book/15-aws-dynamodb-and-sns.md) | aws-sdk-go-v2 Put/Get/Query + Publish |
| 16 | [Feature Flags with Viper](book/16-feature-flags-with-viper.md) | Hot-reloadable flags gated in Fiber |

---

## Hands-On Exercises — 66 Scenarios

Exercises follow an **Exercism-style** format: read the problem, fix the broken code, run tests to validate. Full index: [exercises/README.md](exercises/README.md).

### Book chapters ↔ exercise modules

Chapter numbers and module numbers are **not 1:1**. A single exercise module often covers **several** book chapters (topics grouped for practice). Example: after reading chapters **01–03**, work through **Module 01 — Basics**.

| Exercise module | Book chapters it practices | Exercises |
|-----------------|----------------------------|-----------|
| 01 Basics | **01** Fundamentals, **02** Structs & Methods, **03** Interfaces | 01–08 |
| 02 Collections | **04** Arrays, Slices & Maps | 09–14 |
| 03 Errors | **05** Error Handling | 15–20 |
| 04 Concurrency | **06** Concurrency, **07** Context | 21–30 |
| 05 Testing | **08** Design Patterns, **09** Testing, **10** Profiling | 31–36 |
| 06 BFF / Fiber | **11** Building a BFF with Fiber | 37–42 |
| 07 Resilience | **11** (ops patterns alongside Fiber) | 43–48 |
| 08 Observability | **13** Datadog & Prometheus | 49–53 |
| 09 Viper Config | **12** Viper, **16** Feature Flags | 54–56 |
| 10 Advanced BFF | **11** (aggregation / pooling patterns) | 57–60 |
| 11 Data & Messaging | **14** Redis, **15** DynamoDB & SNS | 61–66 |

### Stack mapping (where to study what)

| Technology | Book chapter | Exercises |
|------------|--------------|-----------|
| Go 1.25+ | all | `exercises/go.mod` |
| Fiber v2 | [11](book/11-building-bff-with-fiber.md) | Module 06 |
| Viper YAML+JSON + fsnotify | [12](book/12-configuration-with-viper.md) | 54–55 |
| Feature flags (Viper) | [16](book/16-feature-flags-with-viper.md) | 56 |
| Datadog dd-trace-go v2 | [13](book/13-observability-datadog-prometheus.md) | 52 |
| Prometheus / fiberprometheus | [13](book/13-observability-datadog-prometheus.md) | 53 |
| Redis go-redis v9 | [14](book/14-redis-with-go-redis.md) | 61–62 |
| DynamoDB aws-sdk-go-v2 | [15](book/15-aws-dynamodb-and-sns.md) | 63–64 |
| SNS aws-sdk-go-v2 | [15](book/15-aws-dynamodb-and-sns.md) | 65–66 |

> **Datadog cost:** the Go client (`dd-trace-go` v2) is **free open source** (Apache-2.0 OR BSD-3-Clause). Sending traces to Datadog’s SaaS is a separate product (free tier available). Exercises use interfaces so you can practice without an agent.
>
> **Internal libs:** company HTTP clients (e.g. quickgo/httpclient) are intentionally omitted — apply the same patterns with `net/http` or Fiber’s agent.

---

## How to Use

1. Read book chapters in order (**00 → 16**)
2. After finishing the chapters listed for a module (see table above), work that exercise module — don’t expect `module-01` to mean “only chapter 01”
3. For each exercise:
   - Read `README.md`
   - Fix the stub in `solution.go`
   - Run `go test -v`
   - Check `_solution/` only if stuck

```bash
go version   # need 1.25+

cd exercises/module-01-basics/01-hello-go
go test -v
```

---

## Prerequisites

- **Go 1.25+** installed
- Basic Java knowledge
- Familiarity with backend / BFF concepts

---

*Based on [Ultimate Go by Ardan Labs](https://github.com/hoanhan101/ultimate-go)*
