# Chapter 05: Error Handling

> No try/catch. No exceptions. Errors are values you handle explicitly.

---

## The Philosophy Shift

### Java — Exceptions

```java
// Java: Exceptions bubble up the call stack automatically
public String readFile(String path) throws IOException {
    // Can throw IOException anywhere
    BufferedReader reader = new BufferedReader(new FileReader(path));
    return reader.readLine();  // May throw
}

// Caller must catch or declare
public void process() {
    try {
        String data = readFile("config.txt");
    } catch (IOException e) {
        logger.error("Failed", e);
    } finally {
        // Always runs
    }
}
```

### Go — Errors as Values

```go
// Go: Errors are returned explicitly
func readFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err  // Return the error
    }
    return string(data), nil
}

// Caller must check
func process() {
    data, err := readFile("config.txt")
    if err != nil {
        log.Error("Failed", "error", err)
        return
    }
    // Use data
}
```

From the ultimate-go resource: *"Integrity matters. Nothing trumps integrity. Error handling is a big part of what we do everyday. It has to be a part of the main code."*

---

## Error Basics

### Creating Errors

```go
import "errors"

// Simple error
err := errors.New("something went wrong")

// Formatted error
err := fmt.Errorf("user %s not found", username)
```

### The error Interface

```go
// Built-in interface — just one method
type error interface {
    Error() string
}

// Any type with Error() string satisfies error
```

### Custom Error Types

```java
// Java: Custom exception class
public class ValidationError extends Exception {
    private String field;
    public ValidationError(String field, String message) {
        super(message);
        this.field = field;
    }
    public String getField() { return field; }
}
```

```go
// Go: Custom error type
type ValidationError struct {
    Field   string
    Message string
}

// Must implement error interface
func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation error on field %s: %s", e.Field, e.Message)
}

// Usage
func validateAge(age int) error {
    if age < 0 {
        return &ValidationError{
            Field:   "age",
            Message: "must be positive",
        }
    }
    return nil
}
```

---

## Error Wrapping (Go 1.13+)

### Adding Context to Errors

```java
// Java: Exception chaining
throw new RuntimeException("Failed to process user", cause);
```

```go
// Go: Wrap errors with %w
func getUser(id int) (*User, error) {
    row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
    var user User
    err := row.Scan(&user.ID, &user.Name)
    if err != nil {
        return nil, fmt.Errorf("getUser(%d): %w", id, err)
    }
    return &user, nil
}
```

### Unwrapping and Inspecting

```go
import "errors"

// errors.Is — checks error chain for specific error
if errors.Is(err, os.ErrNotExist) {
    // File doesn't exist
}

// errors.As — extracts specific error type from chain
var valErr *ValidationError
if errors.As(err, &valErr) {
    fmt.Println("Field:", valErr.Field)
}
```

---

## Sentinel Errors

```go
// Define package-level sentinel errors
var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrRateLimited  = errors.New("rate limited")
)

// Use with errors.Is
func handler(w http.ResponseWriter, r *http.Request) {
    user, err := getUser(1)
    if errors.Is(err, ErrNotFound) {
        http.Error(w, "Not Found", 404)
        return
    }
    if err != nil {
        http.Error(w, "Internal Error", 500)
        return
    }
}
```

---

## Panic and Recover (Go's "Exceptions" — Use Sparingly!)

```java
// Java: RuntimeException for unrecoverable states
throw new RuntimeException("Should never happen");
```

```go
// Go: panic for truly unrecoverable situations
func mustParseConfig(path string) Config {
    data, err := os.ReadFile(path)
    if err != nil {
        panic(fmt.Sprintf("failed to read config: %v", err))
    }
    // ...
}

// Recover (like catch) — only useful in deferred functions
func safeCall() {
    defer func() {
        if r := recover(); r != nil {
            log.Error("Recovered from panic", "error", r)
        }
    }()
    
    mightPanic()
}
```

> ⚠️ **Rule**: Don't use panic/recover for normal error handling. Reserve panics for programmer errors (like index out of bounds) or truly unrecoverable situations.

---

## Error Handling Patterns for High-Traffic Services

### Pattern 1: Always Add Context

```go
// BAD: Lost context
return err

// GOOD: Add call context
return fmt.Errorf("payment service: charge customer %d: %w", customerID, err)
```

### Pattern 2: Handle Once

```go
// BAD: Log AND return (duplicated error in logs)
func process() error {
    err := doSomething()
    if err != nil {
        log.Error("failed", "error", err)  // Logged here
        return err                          // And logged again by caller
    }
}

// GOOD: Either handle OR return
func process() error {
    err := doSomething()
    if err != nil {
        return fmt.Errorf("doSomething: %w", err)  // Caller handles logging
    }
    return nil
}
```

### Pattern 3: Use Custom Types for Domain Errors

```go
type AppError struct {
    Code    string
    Message string
    Err     error
}

func (e *AppError) Error() string { return e.Message }
func (e *AppError) Unwrap() error { return e.Err }

// HTTP handler can inspect error type
func handleError(w http.ResponseWriter, err error) {
    var appErr *AppError
    if errors.As(err, &appErr) {
        switch appErr.Code {
        case "NOT_FOUND":
            w.WriteHeader(http.StatusNotFound)
        case "VALIDATION":
            w.WriteHeader(http.StatusBadRequest)
        default:
            w.WriteHeader(http.StatusInternalServerError)
        }
        json.NewEncoder(w).Encode(map[string]string{"error": appErr.Message})
        return
    }
    w.WriteHeader(http.StatusInternalServerError)
}
```

---

## Key Takeaways

1. **Errors are values** — No try/catch, check return values
2. **Always add context** — Wrap errors with `fmt.Errorf("context: %w", err)`
3. **Use `errors.Is` and `errors.As`** — For inspecting error chains
4. **Handle errors once** — Either log OR return, not both
5. **Custom error types** — For domain-specific error handling
6. **Avoid panic/recover** — Only for programmer errors

---

**Practice**: [Module 03 — Error Handling Exercises →](../exercises/module-03-error-handling/)

Next: [Chapter 06 — Concurrency Model →](06-concurrency-model.md)