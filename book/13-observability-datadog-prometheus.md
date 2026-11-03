# Chapter 13: Observability — Datadog APM & Prometheus

> Production BFFs need traces (where did time go?) and metrics (how much, how fast, how many errors?). Go ecosystems split cleanly: Datadog for hosted APM, Prometheus for metrics you own.

---

## Two Pillars in One Service

| Signal | Question it answers | Typical stack |
|--------|---------------------|---------------|
| **Traces** | Which downstream call failed? What was the critical path? | Datadog APM, OpenTelemetry |
| **Metrics** | What is p99 latency? Error rate? Saturation? | Prometheus + Grafana |

**Java comparison:**

| Go | Java |
|----|------|
| Datadog `dd-trace-go` | Micrometer + Datadog registry, or Spring Cloud Sleuth → Datadog |
| Prometheus client / middleware | Micrometer `PrometheusMeterRegistry` |
| Manual span in code | `@NewSpan`, `Tracer` bean (Sleuth / OTel) |

Spring Boot often exposes `/actuator/prometheus` and auto-instruments JDBC/HTTP. In Go you wire middleware explicitly — fewer surprises, more control.

---

## Datadog `dd-trace-go` v2 (Tracing)

### Library vs backend (important)

- **`github.com/DataDog/dd-trace-go`** is **free and open source** (dual-licensed **Apache-2.0 OR BSD-3-Clause**). You can import it, run tests, and ship binaries without paying Datadog.
- The **Datadog Agent + SaaS backend** is a separate product: free tier exists for small workloads; larger teams use paid plans. Traces are **sent** to an agent (or serverless wrapper); the library does not require a license key to compile.
- Prefer **v2** import paths (`.../v2/...`) for new services. Datadog publishes a [Go version support policy](https://github.com/DataDog/dd-trace-go#support-policy): older tracer majors drop support for EOL Go releases — align CI with **Go 1.25+** as this book targets.

### Setup

```bash
go get github.com/DataDog/dd-trace-go/v2/ddtrace/tracer
go get github.com/DataDog/dd-trace-go/contrib/gofiber/fiber.v2/v2
```

Run the [Datadog Agent](https://docs.datadoghq.com/agent/) locally with APM enabled (`DD_APM_ENABLED=true`), or point to your org’s agent host.

### Manual tracer lifecycle

Unlike Spring’s auto-configured `Tracer`, you start and stop the global tracer in `main`:

```go
package main

import (
    "os"

    "github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
)

func main() {
    tracer.Start(
        tracer.WithService("bff-service"),
        tracer.WithEnv(os.Getenv("DD_ENV")),
        tracer.WithServiceVersion(os.Getenv("DD_VERSION")),
    )
    defer tracer.Stop()

    // ... start Fiber, block until shutdown
}
```

Environment variables commonly used (also set on the Agent):

| Variable | Purpose |
|----------|---------|
| `DD_SERVICE` | Logical service name |
| `DD_ENV` | `prod`, `staging`, … |
| `DD_AGENT_HOST` | Agent hostname (default `localhost`) |
| `DD_TRACE_SAMPLE_RATE` | `0.0`–`1.0` head sampling |

### Fiber v2 middleware (contrib)

```go
import (
    "github.com/gofiber/fiber/v2"
    fibertrace "github.com/DataDog/dd-trace-go/contrib/gofiber/fiber.v2/v2"
)

func main() {
    app := fiber.New()

    app.Use(fibertrace.Middleware(
        fibertrace.WithServiceName("bff-service"),
        fibertrace.WithAnalytics(true),
    ))

    app.Get("/users/:id", getUser)
    // ...
}
```

Each request becomes a span; outbound HTTP clients can be wrapped with contrib modules (`net/http`, `resty`, etc.) to continue the trace.

### Manual spans (business logic)

```go
import (
    "context"

    "github.com/DataDog/dd-trace-go/v2/ddtrace/tracer"
)

func aggregateUserDashboard(ctx context.Context, userID string) error {
    span, ctx := tracer.StartSpanFromContext(ctx, "aggregate.dashboard",
        tracer.ResourceName("user."+userID),
    )
    defer span.Finish()

    if err := fetchOrders(ctx, userID); err != nil {
        span.SetTag("error", err)
        return err
    }
    return fetchPayments(ctx, userID)
}
```

Pass `context.Context` through BFF handlers (Fiber: `c.UserContext()` after middleware attaches the span). This mirrors Sleuth/OTel trace propagation without thread locals.

---

## Prometheus with `fiberprometheus`

Prometheus **pulls** metrics from your process. For Fiber v2, [`github.com/ansrivas/fiberprometheus/v2`](https://github.com/ansrivas/fiberprometheus) exposes standard HTTP metrics and a `/metrics` route.

```bash
go get github.com/ansrivas/fiberprometheus/v2
```

```go
package main

import (
    "github.com/ansrivas/fiberprometheus/v2"
    "github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    prometheus := fiberprometheus.New("bff_service")
    prometheus.RegisterAt(app, "/metrics")
    app.Use(prometheus.Middleware)

    app.Get("/health", func(c *fiber.Ctx) error {
        return c.SendString("ok")
    })

    app.Listen(":8080")
}
```

Scrape `http://bff:8080/metrics` from Prometheus config:

```yaml
scrape_configs:
  - job_name: bff
    static_configs:
      - targets: ["bff:8080"]
```

**Java comparison:** Micrometer registers timers and counters; Prometheus registry exposes text exposition format — same wire format as `fiberprometheus` uses.

### Custom metrics (optional)

For business KPIs (cache hits, aggregation failures), use `prometheus/client_golang` directly and register with the default registry, or extend patterns from exercise 53.

---

## Datadog APM vs Prometheus — Use Both?

They overlap in “observability” but solve different jobs:

| Use Datadog APM when… | Use Prometheus when… |
|----------------------|----------------------|
| You want flame graphs, service maps, trace search out of the box | You run Kubernetes and already have Prometheus/Grafana |
| SREs need correlated logs/traces/RUM in one vendor | You need long-retention custom SLO dashboards on-prem |
| You’re fine sending trace payloads to Datadog Agent | You want pull-based metrics without SaaS dependency |

**Together (common BFF pattern):**

- **Datadog** (or OTel → Datadog): request traces, downstream dependency visibility, latency outliers.
- **Prometheus**: RED metrics (`rate`, `errors`, `duration`) per route, saturation, Go runtime metrics, alert rules in Alertmanager.

Avoid duplicating the same histogram in both systems without reason — pick **one source of truth for alerting** (often Prometheus for infra, Datadog monitors for trace-derived issues).

---

## Fiber: Order of Middleware

```go
tracer.Start(/* ... */)
defer tracer.Stop()

app := fiber.New()

// 1) Tracing first — context for downstream calls
app.Use(fibertrace.Middleware(fibertrace.WithServiceName("bff")))

// 2) Metrics — include all routes except optionally /metrics
prom := fiberprometheus.New("bff")
prom.RegisterAt(app, "/metrics")
app.Use(prom.Middleware)

// 3) Logging, auth, etc.
app.Use(requestLogger)
```

Tracing before metrics ensures failed spans still correlate with access logs.

---

## Local Development Checklist

1. Agent running: `docker run -e DD_APM_ENABLED=true -p 8126:8126 datadog/agent`
2. Export `DD_SERVICE=bff-service`, `DD_ENV=dev`
3. Hit an endpoint; verify trace in Datadog UI (or Agent flare in air-gapped setups)
4. `curl localhost:8080/metrics` — confirm `http_requests_total` increases

---

## Key Takeaways

1. **`dd-trace-go` v2 is OSS** — Apache-2.0 OR BSD-3-Clause; Datadog backend/agent is the commercial piece.
2. **Start/stop the tracer in `main`** — explicit lifecycle like wiring a `Tracer` bean, without Spring magic.
3. **Use Fiber contrib middleware** — `contrib/gofiber/fiber.v2/v2` for automatic HTTP spans on Go 1.25+.
4. **Prometheus via `fiberprometheus/v2`** — `Middleware` + `RegisterAt` for `/metrics` scraping.
5. **Traces for debugging paths; metrics for SLOs** — complementary, not either/or.
6. **Propagate `context.Context`** — BFF aggregation must pass ctx into outbound calls for useful traces.

---

**Practice**: [Module 08 — Observability →](../exercises/module-08-observability/) | [Distributed Tracing →](../exercises/module-08-observability/52-distributed-tracing/) | [Metrics Collection →](../exercises/module-08-observability/53-metrics-collection/)
