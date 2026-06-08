# Chapter 08: Design Patterns

> Go doesn't need Gang of Four patterns. It has its own idiomatic approaches.

---

## Decoupling Through Interfaces

From the ultimate-go resource: *"We want to decouple. The highest level of decoupling we can get is through interfaces."*

```java
// Java: Spring DI with @Autowired
@Service
public class OrderService {
    @Autowired
    private PaymentGateway paymentGateway;
    
    @Autowired
    private NotificationService notificationService;
}
```

```go
// Go: Constructor injection with interfaces
type OrderService struct {
    payment    PaymentGateway
    notify     NotificationService
}

func NewOrderService(p PaymentGateway, n NotificationService) *OrderService {
    return &OrderService{payment: p, notify: n}
}
```

---

## Mocking for Tests

From the ultimate-go resource on mocking:

```go
// Define small interface
type UserFinder interface {
    FindUser(id int) (*User, error)
}

// Real implementation
type DBFinder struct {
    db *sql.DB
}
func (d *DBFinder) FindUser(id int) (*User, error) {
    // Real database query
}

// Mock for testing
type MockFinder struct {
    User  *User
    Error error
}
func (m *MockFinder) FindUser(id int) (*User, error) {
    return m.User, m.Error
}

// Usage in test
func TestGetUser(t *testing.T) {
    mock := &MockFinder{
        User: &User{Name: "Alice"},
    }
    service := NewUserService(mock)
    // Test service without real database
}
```

---

## Type Grouping

```go
// Group related types in a package
package user

type User struct { ID int; Name string }
type Service struct { finder UserFinder }
type Repository interface { FindByID(id int) (*User, error) }

// Factory function returns interface, not concrete type
func NewService(r Repository) *Service {
    return &Service{finder: r}
}
```

---

## Conversion vs Assertion

```go
// Type conversion (between compatible types)
type Celsius float64
type Fahrenheit float64

var c Celsius = 100
var f Fahrenheit = Fahrenheit(c*9/5 + 32)  // Explicit conversion

// Type assertion (extract concrete from interface)
var r io.Reader = &bytes.Buffer{}
b, ok := r.(*bytes.Buffer)  // ok is true if assertion holds
```

---

## Functional Options Pattern (Go's "Builder")

```java
// Java: Builder pattern
Server server = Server.builder()
    .host("localhost")
    .port(8080)
    .timeout(30)
    .build();
```

```go
// Go: Functional options
type Server struct {
    host    string
    port    int
    timeout time.Duration
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) { s.host = host }
}
func WithPort(port int) Option {
    return func(s *Server) { s.port = port }
}
func WithTimeout(d time.Duration) Option {
    return func(s *Server) { s.timeout = d }
}

func NewServer(opts ...Option) *Server {
    s := &Server{host: "0.0.0.0", port: 8080} // defaults
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
server := NewServer(WithHost("localhost"), WithPort(3000))
```

---

## Key Java Patterns → Go Idioms

| Java Pattern | Go Idiom |
|-------------|----------|
| Strategy | Interface + function types |
| Factory | `New*()` functions |
| Builder | Functional options |
| Singleton | Package-level `sync.Once` |
| Observer | Channels |
| Decorator | Middleware chaining |
| Template Method | Interface embedding |
| Null Object | Return zero value or nil |

---

**Practice**: [Module 05 — Testing Exercises →](../exercises/module-05-testing/)

Next: [Chapter 09 — Testing & Benchmarking →](09-testing-and-benchmarking.md)