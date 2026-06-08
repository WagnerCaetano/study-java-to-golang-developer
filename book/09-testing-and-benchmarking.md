# Chapter 09: Testing & Benchmarking

> Go has a built-in testing framework. No JUnit, no Mockito, no Maven Surefire.

---

## Testing Basics

### Java JUnit vs Go Testing

```java
// Java: JUnit 5
@Test
void testAddition() {
    assertEquals(4, calculator.add(2, 2));
}

@Test
void testDivision() {
    assertThrows(ArithmeticException.class, () -> calculator.divide(1, 0));
}
```

```go
// Go: Built-in testing package
func TestAddition(t *testing.T) {
    result := Add(2, 2)
    if result != 4 {
        t.Errorf("Add(2, 2) = %d; want 4", result)
    }
}

func TestDivision(t *testing.T) {
    _, err := Divide(1, 0)
    if err == nil {
        t.Error("expected error for division by zero")
    }
}
```

### Running Tests

```bash
go test                    # All tests in current package
go test ./...              # All tests recursively
go test -v                 # Verbose output
go test -run TestAddition  # Run specific test
go test -race              # With race detector
go test -cover             # Coverage report
go test -coverprofile=cover.out && go tool cover -html=cover.out
```

---

## Table-Driven Tests

From the ultimate-go resource — Go's preferred testing pattern:

```go
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a        int
        b        int
        expected int
    }{
        {"positive numbers", 2, 3, 5},
        {"negative numbers", -1, -2, -3},
        {"zero", 0, 0, 0},
        {"mixed", 5, -3, 2},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
            }
        })
    }
}
```

---

## Sub-Tests with t.Run

```go
func TestUser(t *testing.T) {
    t.Run("create user", func(t *testing.T) {
        u := NewUser("Alice", 30)
        if u.Name != "Alice" {
            t.Error("name mismatch")
        }
    })
    
    t.Run("validate age", func(t *testing.T) {
        _, err := NewUser("Bob", -1)
        if err == nil {
            t.Error("expected error for negative age")
        }
    })
}
```

---

## Mocking with Interfaces

```go
// Interface to mock
type HTTPClient interface {
    Do(req *http.Request) (*http.Response, error)
}

// Mock implementation
type MockClient struct {
    Response *http.Response
    Error    error
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
    return m.Response, m.Error
}

// Test with mock
func TestFetchUser(t *testing.T) {
    mockResp := &http.Response{
        Body: io.NopCloser(strings.NewReader(`{"name":"Alice"}`)),
    }
    client := &MockClient{Response: mockResp}
    
    service := NewUserService(client)
    user, err := service.FetchUser("1")
    
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if user.Name != "Alice" {
        t.Errorf("expected Alice, got %s", user.Name)
    }
}
```

---

## HTTP Handler Testing

```go
func TestHandler(t *testing.T) {
    req := httptest.NewRequest("GET", "/users/1", nil)
    rec := httptest.NewRecorder()
    
    handler := GetUserHandler(mockService)
    handler.ServeHTTP(rec, req)
    
    if rec.Code != http.StatusOK {
        t.Errorf("expected 200, got %d", rec.Code)
    }
}
```

---

## Benchmarking

```go
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(2, 3)
    }
}

func BenchmarkLargeSlice(b *testing.B) {
    data := make([]int, 1000)
    b.ResetTimer()  // Don't count setup
    
    for i := 0; i < b.N; i++ {
        process(data)
    }
}
```

```bash
go test -bench=. -benchmem
# Output: BenchmarkAdd-8   1000000000   0.256 ns/op   0 B/op   0 allocs/op
```

---

## Key Takeaways

1. **No test framework needed** — `go test` is built-in
2. **Table-driven tests** — Go's canonical testing pattern
3. **Mock with interfaces** — No Mockito needed
4. **`httptest` package** — For HTTP handler testing
5. **Benchmarks are first-class** — `go test -bench`
6. **Always run with `-race`** — Catch data races early

---

**Practice**: [Module 05 — Testing Exercises →](../exercises/module-05-testing/)

Next: [Chapter 10 — Profiling & Optimization →](10-profiling-and-optimization.md)