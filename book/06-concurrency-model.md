# Chapter 06: Concurrency Model

> Goroutines and channels: the reason many Java teams switch to Go.

---

## The Mental Model Shift

From the ultimate-go resource: *"Concurrency is about managing a lot of things at once. Parallelism is about doing a lot of things at once."*

### Java Threads vs Go Goroutines

| Aspect | Java Thread | Go Goroutine |
|---------|------------|--------------|
| Stack size | ~1 MB (fixed) | ~2 KB (grows as needed) |
| Creation cost | Expensive (OS thread) | Cheap (user-space) |
| Scheduling | OS preemptive scheduler | Go cooperative scheduler |
| Context switch | ~1-10 μs (kernel) | ~100 ns (user-space) |
| Practical limit | ~1,000-10,000 | ~100,000+ |

---

## Goroutines

### Launching Goroutines

```java
// Java: ExecutorService
ExecutorService executor = Executors.newFixedThreadPool(10);
executor.submit(() -> doWork());
executor.shutdown();
```

```go
// Go: Just add 'go' keyword
go doWork()                    // Simple function
go func() {                    // Anonymous function
    fmt.Println("working")
}()
go func(name string) {         // With parameters
    fmt.Println("hello", name)
}("Alice")
```

### The Scheduler (from ultimate-go)

Go's scheduler uses:
- **G** — Goroutine
- **M** — Machine (OS thread)
- **P** — Logical Processor (determined by GOMAXPROCS)

```
      GRQ (Global Run Queue)
    -------------------------
   |  G1  |  G2  |  G3  | ...
    -------------------------

      P1                P2
    ------            ------
   | M1   |          | M2   |
    ------            ------
   | LRQ  |          | LRQ  |
   | G4   |          | G6   |
   | G5   |          | G7   |
    ------            ------
```

Key scheduling decision points:
1. `go` keyword (goroutine creation)
2. System calls
3. Channel operations (blocking/unblocking)
4. Garbage collection

---

## Channels

From the ultimate-go resource: *"Channels are for orchestration. Think of channels as a way of signaling events, not just queues."*

### Unbuffered Channels — Guaranteed Delivery

```go
// Signal WITH data — guaranteed the signal was received
ch := make(chan string)

go func() {
    ch <- "hello"  // Blocks until someone receives
}()

msg := <-ch  // Blocks until data arrives
fmt.Println(msg)  // "hello"
```

### Buffered Channels — Lower Latency, No Guarantee

```go
// Signal with data — send doesn't block (until buffer is full)
ch := make(chan string, 1)  // Buffer of 1

go func() {
    ch <- "hello"  // Doesn't block — buffer has space
    // Goroutine can continue immediately
}()

msg := <-ch  // Receives when ready
```

### Signaling Without Data

```go
// Close-based signaling (no data needed)
done := make(chan struct{})

go func() {
    defer close(done)  // Signal completion
    doWork()
}()

<-done  // Wait for signal
fmt.Println("work done")
```

### Channel Direction (Documentation)

```go
// Send-only channel
func producer(ch chan<- int) {
    ch <- 42
}

// Receive-only channel
func consumer(ch <-chan int) {
    val := <-ch
    fmt.Println(val)
}

// Bidirectional channel
func pipe(ch chan int) {
    ch <- <-ch  // Receive and send
}
```

---

## Select — Multiplexing Channels

```java
// Java: No direct equivalent — would use CompletableFuture.anyOf()
CompletableFuture<Object> result = CompletableFuture.anyOf(future1, future2);
```

```go
// Go: Select waits on multiple channel operations
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case <-time.After(5 * time.Second):
    fmt.Println("Timeout!")
default:
    fmt.Println("No channel ready")  // Non-blocking
}
```

### Timeout Pattern

```go
func fetchWithTimeout(url string) (string, error) {
    result := make(chan string, 1)
    
    go func() {
        resp, _ := http.Get(url)
        defer resp.Body.Close()
        body, _ := io.ReadAll(resp.Body)
        result <- string(body)
    }()
    
    select {
    case data := <-result:
        return data, nil
    case <-time.After(5 * time.Second):
        return "", fmt.Errorf("timeout fetching %s", url)
    }
}
```

---

## Common Concurrency Patterns

### Fan-Out / Fan-In

```go
func fanOutFanIn(input <-chan int, workers int) <-chan int {
    results := make(chan int)
    
    // Fan-out: launch multiple workers
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for val := range input {
                results <- process(val)
            }
        }()
    }
    
    // Fan-in: collect results
    go func() {
        wg.Wait()
        close(results)
    }()
    
    return results
}
```

### Pipeline

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

// Pipeline: generate → square → consume
for result := range square(square(generate(1, 2, 3))) {
    fmt.Println(result)  // 1, 16, 81
}
```

### Worker Pool

```go
func workerPool(jobs <-chan Job, results chan<- Result, workers int) {
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            for job := range jobs {
                results <- processJob(id, job)
            }
        }(i)
    }
    go func() {
        wg.Wait()
        close(results)
    }()
}
```

---

## Data Races and Mutexes

### Detecting Data Races

```bash
go test -race ./...   # Built-in race detector!
```

```go
// RACE CONDITION: Multiple goroutines accessing shared data
var counter int
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        counter++  // DATA RACE!
        wg.Done()
    }()
}
wg.Wait()
```

### Fix with Mutex

```go
var counter int
var mu sync.Mutex
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        mu.Lock()
        counter++  // Safe
        mu.Unlock()
        wg.Done()
    }()
}
```

### Fix with Atomic

```go
var counter int64
var wg sync.WaitGroup

for i := 0; i < 1000; i++ {
    wg.Add(1)
    go func() {
        atomic.AddInt64(&counter, 1)  // Lock-free
        wg.Done()
    }()
}
```

### Prefer Channels Over Mutexes

From the Go proverb: *"Don't communicate by sharing memory; share memory by communicating."*

```go
// Instead of mutex-protected map:
// var mu sync.Mutex
// m := make(map[string]int)

// Use a goroutine with channel access:
type command struct {
    key    string
    result chan int
}

func safeMap(cmds <-chan command) {
    m := make(map[string]int)  // Only this goroutine accesses it
    for cmd := range cmds {
        cmd.result <- m[cmd.key]
    }
}
```

---

## Key Takeaways

1. **Goroutines are cheap** — Launch thousands without worry
2. **Channels are for signaling** — Not just data queues
3. **Use `select` for multiplexing** — Handle multiple channels
4. **Always use the race detector** — `go test -race`
5. **Prefer channels** over shared memory + mutexes
6. **Close channels to signal completion** — Don't send sentinel values

---

**Practice**: [Module 04 — Concurrency Exercises →](../exercises/module-04-concurrency/)

Next: [Chapter 07 — Context & Cancellation →](07-context-and-cancellation.md)