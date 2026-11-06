# Chapter 00: Introduction — Why Go for a Java Developer

> "Go is not meant to be a better Java. Go is meant to be a better C for the modern world."  
> But for backend developers, Go fills the gap Java leaves in high-traffic, low-latency scenarios.

---

## 🤔 Why Should a Java Developer Learn Go?

As a Java developer, you already understand:
- Object-oriented programming
- Design patterns (Factory, Strategy, Observer, etc.)
- Dependency injection
- Enterprise backend development
- JVM internals (garbage collection, classloading, JIT compilation)

**So why learn Go?**

| Aspect | Java | Go |
|--------|------|----|
| Startup time | Seconds (JVM warmup) | Milliseconds (compiled binary) |
| Memory footprint | Hundreds of MB baseline | ~10 MB baseline |
| Concurrency model | Threads + CompletableFuture | Goroutines + Channels |
| Deployment | JAR + JVM required | Single binary |
| Build time | Can be slow | Seconds |
| Garbage collection | Complex, tunable | Simple, low-latency |
| Binary size | Large (JVM overhead) | Small, self-contained |
| Kubernetes readiness | Needs JVM tuning | Native, lightweight |

---

## 🔑 Key Philosophical Differences

### 1. Simplicity Over Complexity
Java embraces complexity as a tool for abstraction (generics, inheritance hierarchies, design patterns). Go deliberately rejects complexity:

```java
// Java: Strategy pattern with interface + multiple implementations
public interface PaymentStrategy {
    void pay(BigDecimal amount);
}
public class CreditCardPayment implements PaymentStrategy { ... }
public class PixPayment implements PaymentStrategy { ... }
```

```go
// Go: A function type is often enough
type PaymentFunc func(amount float64) error

payWithCredit := func(amount float64) error { return nil }
payWithPix := func(amount float64) error { return nil }
```

### 2. Composition Over Inheritance
Java uses class inheritance (`extends`, `implements`). Go has **no class inheritance** — only composition through struct embedding and interfaces.

```java
// Java: Inheritance hierarchy
public class Animal { }
public class Dog extends Animal { }
public class GuideDog extends Dog { }
```

```go
// Go: Composition
type Animal struct{ Name string }
type Dog struct{ Animal }                    // Embedding
type GuideDog struct{ Dog; Owner string }   // More embedding
```

### 3. Errors Are Values, Not Exceptions
Java uses exceptions (`try/catch/finally`). Go returns errors as values:

```java
// Java: Exceptions
try {
    String data = readFile("config.txt");
} catch (IOException e) {
    logger.error("Failed to read", e);
} finally {
    cleanup();
}
```

```go
// Go: Errors as values
data, err := readFile("config.txt")
if err != nil {
    log.Error("Failed to read", "error", err)
    return err
}
```

### 4. Goroutines vs Java Threads
A Java thread costs ~1MB of stack memory. A goroutine starts with ~2KB:

```java
// Java: Thread-per-request model (expensive)
ExecutorService pool = Executors.newFixedThreadPool(200);
for (Request req : requests) {
    pool.submit(() -> processRequest(req));
}
```

```go
// Go: Goroutines are cheap
for _, req := range requests {
    go processRequest(req)  // 2KB per goroutine!
}
```

### 5. No Null — But Be Careful
Java has `null` and `NullPointerException`. Go has `nil` for pointers, slices, maps, channels, interfaces, and functions. However, Go does **not** have optional types or null-safety like Kotlin.

---

## 🛠 Tooling Setup

### Installing Go
```bash
# macOS
brew install go

# Linux
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz

# Windows
# Download from https://go.dev/dl/
```

### Essential Commands (Java → Go equivalents)

| Task | Java | Go |
|------|------|----|
| Build | `mvn compile` | `go build` |
| Test | `mvn test` | `go test` |
| Run | `java -jar app.jar` | `go run main.go` |
| Dependencies | `pom.xml` / `build.gradle` | `go.mod` |
| Format | Checkstyle / Spotless | `go fmt` / `goimports` |
| Lint | SonarQube / PMD | `golangci-lint` |
| Vet | SpotBugs | `go vet` |
| Documentation | Javadoc | `go doc` |

### IDE Setup
- **VS Code** with the official Go extension (recommended)
- **GoLand** by JetBrains (familiar for IntelliJ users)

### Your First Go Project
```bash
mkdir myproject && cd myproject
go mod init github.com/yourname/myproject
```

This creates a `go.mod` file (similar to `pom.xml`):

```
module github.com/yourname/myproject

go 1.25
```

This book targets **Go 1.25+**.

---

## 📚 Study Flow

This book follows a deliberate progression:

1. **[Language Fundamentals](01-language-fundamentals.md)** — Learn the basic building blocks
2. **[Structs & Methods](02-structs-and-methods.md)** — Understand Go's take on "objects"
3. **[Interfaces & Polymorphism](03-interfaces-and-polymorphism.md)** — Go's implicit contracts
4. **[Collections](04-arrays-slices-maps.md)** — Slices and maps (no ArrayList, no HashMap)
5. **[Error Handling](05-error-handling.md)** — No exceptions, no excuses
6. **[Concurrency](06-concurrency-model.md)** — Goroutines and channels
7. **[Context](07-context-and-cancellation.md)** — Request-scoped values and cancellation
8. **[Design Patterns](08-design-patterns.md)** — Go-native patterns
9. **[Testing](09-testing-and-benchmarking.md)** — Built-in testing framework
10. **[Profiling](10-profiling-and-optimization.md)** — Performance optimization
11. **[BFF with Fiber](11-building-bff-with-fiber.md)** — Fiber v2 web framework
12. **[Configuration with Viper](12-configuration-with-viper.md)** — YAML/JSON + hot reload
13. **[Observability](13-observability-datadog-prometheus.md)** — Datadog APM + Prometheus
14. **[Redis](14-redis-with-go-redis.md)** — Caching with go-redis v9
15. **[AWS DynamoDB & SNS](15-aws-dynamodb-and-sns.md)** — Persistence and messaging
16. **[Feature Flags](16-feature-flags-with-viper.md)** — Viper-driven flags

Each chapter links to hands-on exercises where you'll practice what you've learned.

### Chapters vs exercise modules (important)

**Book chapter numbers and exercise module numbers are not the same thing.**  
One exercise module can cover **more than one** book chapter. Topics are grouped so you practice related ideas together.

| After reading book chapters… | Do exercise module… |
|------------------------------|---------------------|
| 01 + 02 + 03 (fundamentals, structs, interfaces) | [Module 01 — Basics](../exercises/module-01-basics/) |
| 04 (collections) | [Module 02 — Collections](../exercises/module-02-collections/) |
| 05 (errors) | [Module 03 — Error Handling](../exercises/module-03-error-handling/) |
| 06 + 07 (concurrency, context) | [Module 04 — Concurrency](../exercises/module-04-concurrency/) |
| 08 + 09 + 10 (patterns, testing, profiling) | [Module 05 — Testing](../exercises/module-05-testing/) |
| 11 (Fiber BFF) | [Modules 06, 07, 10](../exercises/README.md) |
| 12 + 16 (Viper, feature flags) | [Module 09 — Viper Config](../exercises/module-09-viper-config/) |
| 13 (observability) | [Module 08 — Observability](../exercises/module-08-observability/) |
| 14 + 15 (Redis, DynamoDB, SNS) | [Module 11 — Data & Messaging](../exercises/module-11-data-messaging/) |

Full mapping: [root README](../README.md) and [exercises/README.md](../exercises/README.md).

---

## 🧠 Mental Model Shift

As a Java developer learning Go, the most important shift is:

> **Stop thinking in classes and start thinking in functions and data.**

Java organizes code around classes (nouns). Go organizes code around packages and functions (verbs). A typical Go program has:
- **Data types** (structs) that are plain data holders
- **Functions** that operate on data
- **Interfaces** that define behavior contracts
- **Packages** that group related functionality

There are no constructors, no `this`, no `super`, no generics with bounded types (until recently), no checked exceptions, no annotations.

**Embrace the simplicity. It's a feature, not a limitation.**

---

Next: [Chapter 01 — Language Fundamentals →](01-language-fundamentals.md)