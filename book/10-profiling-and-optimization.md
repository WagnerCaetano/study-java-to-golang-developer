# Chapter 10: Profiling & Optimization

> In high-traffic scenarios, every allocation matters. Learn to profile before you optimize.

---

## Escape Analysis

From the ultimate-go resource: The Go compiler decides what goes on the stack (free) vs the heap (GC pressure).

```bash
# Check escape analysis
go build -gcflags="-m" main.go
```

```go
// Stack allocation (good — no GC pressure)
func sum(a, b int) int {
    result := a + b  // Stays on stack
    return result
}

// Heap allocation (escaped — GC must manage)
func createUser() *User {
    u := User{Name: "Alice"}  // Escapes to heap (returned as pointer)
    return &u
}
```

### Common Escape Triggers
- Returning pointer to local variable
- Storing reference in interface
- Closure capturing variable
- Slice/map that grows beyond stack size

---

## Memory Tracing

```go
func printMemStats() {
    var m runtime.MemStats
    runtime.ReadMemStats(&m)
    fmt.Printf("Alloc = %v KB\n", m.Alloc/1024)
    fmt.Printf("TotalAlloc = %v KB\n", m.TotalAlloc/1024)
    fmt.Printf("Sys = %v KB\n", m.Sys/1024)
    fmt.Printf("NumGC = %v\n", m.NumGC)
}
```

---

## Stack Traces

```go
// Use runtime.Caller for debugging
func trace() {
    pc, file, line, ok := runtime.Caller(1)
    if ok {
        fn := runtime.FuncForPC(pc)
        fmt.Printf("%s:%d %s\n", file, line, fn.Name())
    }
}
```

---

## CPU Profiling

```go
import "runtime/pprof"

func main() {
    f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // Your code here
}
```

```bash
go tool pprof cpu.prof
# Interactive: top, list FuncName, web
```

## HTTP Profiling Endpoint

```go
import _ "net/http/pprof"

go http.ListenAndServe("localhost:6060", nil)
```

```bash
go tool pprof http://localhost:6060/debug/pprof/profile
```

---

## Optimization Tips for High-Traffic

1. **Pre-allocate slices** — `make([]T, 0, estimatedSize)`
2. **Reuse buffers with sync.Pool** — Avoid repeated allocations
3. **Use value types when possible** — Avoid pointer indirection
4. **Minimize interface conversions** — Each has a cost
5. **Profile before optimizing** — Don't guess, measure

---

**Practice**: [Module 05 — Benchmarking Exercise →](../exercises/module-05-testing/35-benchmarking/)

Next: [Chapter 11 — Building a BFF with Fiber →](11-building-bff-with-fiber.md)