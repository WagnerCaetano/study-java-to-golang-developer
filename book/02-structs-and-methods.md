# Chapter 02: Structs & Methods

> Go doesn't have classes. It has structs with methods. Learn how this changes everything.

---

## Structs — Go's "Classes"

### Java Class vs Go Struct

```java
// Java: Class with constructor, fields, methods
public class User {
    private String name;
    private int age;
    private String email;
    
    public User(String name, int age, String email) {
        this.name = name;
        this.age = age;
        this.email = email;
    }
    
    public String getName() { return name; }
    public int getAge() { return age; }
}
```

```go
// Go: Struct — just data, no behavior attached
type User struct {
    Name  string  // exported (public)
    Age   int     // exported (public)
    email string  // unexported (package-private)
}
```

### Creating Struct Instances

```go
// Zero value (all fields initialized to their zero values)
var u1 User  // Name: "", Age: 0, email: ""

// Using field names (recommended)
u2 := User{
    Name:  "Alice",
    Age:   30,
    email: "alice@example.com",
}

// Positional (fragile — avoid if struct might change)
u3 := User{"Bob", 25, "bob@example.com"}

// Using new (returns pointer, zero-valued)
u4 := new(User)  // *User, all fields zero

// Using & (returns pointer, can initialize)
u5 := &User{Name: "Charlie", Age: 35}
```

### No Constructors — Use Factory Functions

```java
// Java: Constructor overloading
public User(String name) { this(name, 0, ""); }
public User(String name, int age) { this(name, age, ""); }
public User(String name, int age, String email) { ... }
```

```go
// Go: Factory functions (convention: NewTypeName)
func NewUser(name string, age int, email string) *User {
    return &User{
        Name:  name,
        Age:   age,
        email: email,
    }
}

// With validation
func NewUser(name string, age int) (*User, error) {
    if name == "" {
        return nil, errors.New("name is required")
    }
    if age < 0 {
        return nil, errors.New("age must be positive")
    }
    return &User{Name: name, Age: age}, nil
}
```

---

## Methods — Attaching Behavior to Data

### Java Methods vs Go Receiver Functions

```java
// Java: Methods live inside the class
public class User {
    private String name;
    
    public String Greet() {
        return "Hello, " + name;
    }
}
```

```go
// Go: Methods are functions with a "receiver" parameter
func (u User) Greet() string {
    return "Hello, " + u.Name
}

// Usage is the same:
user := User{Name: "Alice"}
fmt.Println(user.Greet())  // "Hello, Alice"
```

### Value Receivers vs Pointer Receivers

This is **critical** and has no direct Java equivalent:

```go
// VALUE RECEIVER: Works on a COPY of the struct
func (u User) GetName() string {
    return u.Name  // Reads from a copy
}

// POINTER RECEIVER: Works on the ORIGINAL struct
func (u *User) SetName(name string) {
    u.Name = name  // Modifies the original
}
```

```go
user := User{Name: "Alice"}

user.SetName("Bob")       // OK — modifies original
fmt.Println(user.Name)    // "Bob"

name := user.GetName()    // OK — reads from copy (safe)
```

### When to Use Which?

| Scenario | Receiver Type | Why |
|----------|--------------|-----|
| Need to mutate the struct | `*T` (pointer) | Must modify original |
| Struct is large (>64 bytes) | `*T` (pointer) | Avoid copying |
| Need consistency | `*T` (pointer) | If some methods need pointer, all should use it |
| Small struct, read-only | `T` (value) | Safe, no mutation possible |
| Primitive-like types | `T` (value) | time.Time, big.Int exceptions |

> 🔑 **Best Practice**: When in doubt, use **pointer receivers**. They're safer for mutation and avoid copies.

### Go's "this" is the Receiver

```java
// Java
public class Counter {
    private int count;
    
    public void increment() {
        this.count++;       // 'this' is implicit
    }
}
```

```go
// Go: Choose your own receiver name (convention: 1-2 letter abbreviation)
func (c *Counter) Increment() {
    c.Count++  // c is explicit — there is no 'this'
}

// Common conventions:
func (u *User) Save()    { ... }  // 'u' for User
func (s *Server) Start() { ... }  // 's' for Server
func (h *Handler) Serve() { ... } // 'h' for Handler
```

---

## Pointers — Understanding Go's Memory Model

### Java References vs Go Pointers

```java
// Java: All objects are reference types
User u1 = new User("Alice");
User u2 = u1;      // u2 points to same object
u2.setName("Bob");
System.out.println(u1.getName()); // "Bob" — shared!
```

```go
// Go: Value types vs Pointer types
u1 := User{Name: "Alice"}
u2 := u1           // COPY — different memory
u2.Name = "Bob"
fmt.Println(u1.Name)  // "Alice" — original unchanged!

// With pointers (similar to Java references)
u3 := &User{Name: "Alice"}  // *User
u4 := u3                     // Both point to same memory
u4.Name = "Bob"
fmt.Println(u3.Name)  // "Bob" — shared!
```

### Pointer Mechanics

```go
x := 42
p := &x         // p is *int, points to x
fmt.Println(*p) // 42 — dereference
*p = 100        // Modify through pointer
fmt.Println(x)  // 100

// Go pointers are simpler than C/C++
// - No pointer arithmetic
// - No -> operator (use . for everything)
```

### Escape Analysis and the Stack vs Heap

In Java, all objects go on the heap. Go's compiler performs **escape analysis** to decide:

```go
func createUser() *User {
    u := User{Name: "Alice"}
    return &u  // 'u' escapes to heap (returned as pointer)
}

func useUser() User {
    u := User{Name: "Bob"}
    return u  // 'u' stays on stack (returned as value)
}
```

Check with `go build -gcflags="-m"`:
```
createUser: u escapes to heap
useUser: u does not escape
```

> 🎯 **High-Traffic Insight**: Minimize heap allocations. Every heap allocation is work for the GC. In high-traffic services, stack-allocated data is essentially free.

---

## Struct Embedding — Go's "Inheritance"

### Java Inheritance vs Go Embedding

```java
// Java: Inheritance
public class BaseEntity {
    private Long id;
    private LocalDateTime createdAt;
}

public class User extends BaseEntity {
    private String name;
}
```

```go
// Go: Embedding (composition, NOT inheritance)
type BaseEntity struct {
    ID        int64
    CreatedAt time.Time
}

type User struct {
    BaseEntity       // Embedded (no field name)
    Name     string
}

// Usage
u := User{
    BaseEntity: BaseEntity{ID: 1, CreatedAt: time.Now()},
    Name:       "Alice",
}
fmt.Println(u.ID)  // Access embedded field directly
```

### Embedding is NOT Inheritance

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return "..."
}

type Dog struct {
    Animal
    Breed string
}

func (d Dog) Speak() string {
    return "Woof!"
}

d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
d.Speak()       // "Woof!" — Dog's method wins
d.Animal.Speak() // "..." — Can still access Animal's method
```

> ⚠️ Key difference: In Java, `super.speak()` calls the parent. In Go, the outer type's method shadows the inner type's method. There is no `super`.

---

## Memory Layout and Padding

From the ultimate-go resource, struct field ordering matters:

```go
// BAD: 24 bytes (on 64-bit) due to padding
type Bad struct {
    flag    bool    // 1 byte + 7 padding
    counter int64   // 8 bytes
    pi      float32 // 4 bytes + 4 padding
}

// GOOD: 16 bytes — order from largest to smallest
type Good struct {
    counter int64   // 8 bytes
    pi      float32 // 4 bytes
    flag    bool    // 1 byte + 3 padding
}
```

> 🎯 **High-Traffic Insight**: In data-intensive applications with millions of structs, proper field ordering can reduce memory usage by 30-50%.

---

## Key Takeaways

1. **Structs are just data** — No behavior, no inheritance, no constructors
2. **Use factory functions** (`New*`) instead of constructors
3. **Pointer receivers for mutation**, value receivers for read-only
4. **Embedding is composition**, not inheritance — there is no `super`
5. **Field ordering matters** — Lay out from largest to smallest to minimize padding
6. **Minimize heap allocations** — Use value types when possible in high-traffic code

---

**Practice**: [Module 01 Exercises →](../exercises/module-01-basics/)

Next: [Chapter 03 — Interfaces & Polymorphism →](03-interfaces-and-polymorphism.md)