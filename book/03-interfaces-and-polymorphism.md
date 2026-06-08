# Chapter 03: Interfaces & Polymorphism

> Go interfaces are implicit. If it walks like a duck and talks like a duck, it's a duck.

---

## The Big Difference: Implicit Implementation

### Java — Explicit `implements`

```java
// Java: You MUST declare that you implement the interface
public interface Reader {
    int read(byte[] b) throws IOException;
}

public class FileReader implements Reader {
    @Override
    public int read(byte[] b) throws IOException {
        // implementation
    }
}
```

### Go — Implicit Satisfaction

```go
// Go: Just implement the method — no 'implements' keyword
type Reader interface {
    Read(b []byte) (int, error)
}

type File struct {
    Name string
}

// File automatically satisfies Reader — no declaration needed
func (f File) Read(b []byte) (int, error) {
    data := []byte("file data")
    copy(b, data)
    return len(data), nil
}
```

> 🔑 **Key Insight**: In Go, you can satisfy an interface **without knowing it exists**. This enables decoupled, bottom-up design.

---

## Interface Mechanics

### Interface Is a Two-Word Data Structure

From the ultimate-go resource, an interface value has two words:

```
     Interface
  ---------------
 | *iTable ptr   | --> iTable { type info, method pointers }
  ---------------
 | *data ptr     | --> Copy of the concrete value
  ---------------
```

```go
var r Reader           // nil interface (both words nil)
r = File{Name: "data.json"}  // Now r holds a File
```

### Nil Interface vs Nil Value Under Interface

This is a **common gotcha** for Java developers:

```go
var f *File = nil   // f is nil
var r Reader = f    // r is NOT nil! (has type info)

// r != nil — because the interface has type information
// This causes confusing bugs!

// Correct nil check:
if r != nil {
    // This executes even though f is nil!
}
```

> ⚠️ **Rule**: Never return a concrete nil typed as an interface. Return `nil` directly.

```go
// BAD: Returns non-nil interface with nil value
func getReader() Reader {
    var f *File = nil
    return f  // interface is not nil!
}

// GOOD: Return nil directly
func getReader() Reader {
    return nil  // interface is truly nil
}
```

---

## Interface Composition

### Java Interface Inheritance vs Go Embedding

```java
// Java: Interface inheritance
public interface Reader {
    int read(byte[] b);
}
public interface Writer {
    int write(byte[] b);
}
public interface ReadWriter extends Reader, Writer {
}
```

```go
// Go: Interface embedding (composition)
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReadWriter interface {
    Reader
    Writer
}
```

### The Empty Interface — `interface{}` / `any`

```java
// Java: Object is the root of all types
Object obj = "hello";
Object obj2 = 42;
```

```go
// Go: interface{} (or 'any' in Go 1.18+) accepts any type
var obj any = "hello"
var obj2 any = 42

// But avoid it — Go prefers explicit types
// Use when you truly need dynamic typing (JSON parsing, etc.)
```

---

## Polymorphism in Practice

### The Power of Decoupling

```go
// Define behavior
type Serializer interface {
    Serialize() ([]byte, error)
}

// Multiple implementations
type JSONSerializer struct {
    Data any
}

func (j JSONSerializer) Serialize() ([]byte, error) {
    return json.Marshal(j.Data)
}

type ProtobufSerializer struct {
    Message proto.Message
}

func (p ProtobufSerializer) Serialize() ([]byte, error) {
    return proto.Marshal(p.Message)
}

// Polymorphic function — doesn't know about concrete types
func sendToQueue(s Serializer) error {
    data, err := s.Serialize()
    if err != nil {
        return err
    }
    // send data...
    return nil
}
```

### Type Assertion and Type Switch

```java
// Java: instanceof + casting
if (animal instanceof Dog) {
    Dog dog = (Dog) animal;
    dog.fetch();
}
```

```go
// Go: Type assertion
var r Reader = File{Name: "data.json"}
f := r.(File)         // Direct assertion — panics if wrong type
f, ok := r.(File)     // Safe assertion — ok is false if wrong type

// Type switch (much cleaner than Java's chained instanceof)
switch v := r.(type) {
case File:
    fmt.Println("File:", v.Name)
case *bytes.Buffer:
    fmt.Println("Buffer:", v.Len())
default:
    fmt.Println("Unknown type")
}
```

---

## Interface Design Rules

### Rule 1: Keep Interfaces Small

```go
// BAD: Fat interface (Java style)
type UserRepository interface {
    FindByID(id int) (*User, error)
    FindAll() ([]User, error)
    Create(user *User) error
    Update(user *User) error
    Delete(id int) error
    FindByEmail(email string) (*User, error)
    FindByRole(role string) ([]User, error)
}
```

```go
// GOOD: Small, focused interfaces (Go style)
type UserFinder interface {
    FindByID(id int) (*User, error)
}

type UserCreator interface {
    Create(user *User) error
}

type UserDeleter interface {
    Delete(id int) error
}
```

### Rule 2: Accept Interfaces, Return Structs

```go
// Function accepts interface for flexibility
func ProcessData(r io.Reader) error {
    // Works with any Reader: files, network, buffers, etc.
    data, err := io.ReadAll(r)
    // ...
}

// Function returns concrete type for clarity
func NewFileReader(path string) *os.File {
    f, _ := os.Open(path)
    return f
}
```

### Rule 3: Define Interfaces Where They're Used

```java
// Java: Define interfaces in their own package
// com.example.contracts.Reader
```

```go
// Go: Define the interface in the package that USES it, not provides it
// In package "processor":
type Reader interface {
    Read(p []byte) (int, error)
}

func Process(r Reader) { ... }

// In package "file" (no import of processor needed!):
type File struct { ... }
func (f File) Read(p []byte) (int, error) { ... }
```

---

## Common Standard Library Interfaces

| Interface | Methods | Used By |
|-----------|---------|---------|
| `io.Reader` | `Read(p []byte) (n int, err error)` | Files, network, buffers |
| `io.Writer` | `Write(p []byte) (n int, err error)` | Files, network, buffers |
| `fmt.Stringer` | `String() string` | fmt.Println formatting |
| `error` | `Error() string` | All error handling |
| `sort.Interface` | `Len(), Less(), Swap()` | Sorting slices |
| `http.Handler` | `ServeHTTP(w, r)` | HTTP routing |

---

## Key Takeaways

1. **Interfaces are implicit** — No `implements` keyword
2. **Keep interfaces small** — 1-2 methods is ideal
3. **Accept interfaces, return structs** — Flexibility at boundaries
4. **Define interfaces where used** — Not where implemented
5. **Beware nil interface trap** — A nil value inside a non-nil interface
6. **Use type switches** — Not chained type assertions

---

**Practice**: [Module 01 Exercise 08 — Interfaces →](../exercises/module-01-basics/08-interfaces/)

Next: [Chapter 04 — Arrays, Slices & Maps →](04-arrays-slices-maps.md)