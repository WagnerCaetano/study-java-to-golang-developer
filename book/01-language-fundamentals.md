# Chapter 01: Language Fundamentals

> From Java's verbose type system to Go's pragmatic simplicity.

---

## Variables

### Java vs Go Variable Declaration

```java
// Java - Explicit type declaration
String name = "Go Developer";
int age = 30;
double price = 19.99;
boolean active = true;
```

```go
// Go - Type inference with :=
name := "Go Developer"     // string
age := 30                  // int
price := 19.99             // float64
active := true             // bool
```

Go also supports explicit type declaration (use when the zero value matters):

```go
var name string = "Go Developer"
var age int = 30
var price float64 = 19.99
var active bool = true
```

### Zero Values (No `null` for basic types!)

In Java, uninitialized fields are `null` (for objects) or `0`/`false` (for primitives). In Go, **every type has a zero value**:

```go
var s string   // "" (empty string)
var i int      // 0
var f float64  // 0.0
var b bool     // false
var p *int     // nil (pointers can be nil)
```

### Short Variable Declaration (`:=`)

The `:=` operator declares AND assigns. It can only be used **inside functions**:

```go
func main() {
    count := 10        // OK: declare and assign
    count = 20         // OK: just assign (no colon)
    // count := 30     // ERROR: no new variables on left side of :=
    
    name, age := "Alice", 25  // Multiple assignment
}
```

### Constants

```java
// Java
public static final String VERSION = "1.0";
public static final int MAX_RETRIES = 3;
```

```go
// Go
const version = "1.0"
const maxRetries = 3

// Grouped constants
const (
    StatusOK    = 200
    StatusError = 500
)
```

### iota — Go's Enum Generator

Go doesn't have Java's `enum`. Instead, it uses `iota`:

```java
// Java
public enum HttpStatus {
    OK(200), NOT_FOUND(404), ERROR(500);
    private final int code;
    HttpStatus(int code) { this.code = code; }
}
```

```go
// Go
type HttpStatus int

const (
    StatusOK HttpStatus = iota + 200  // 200
    StatusNotFound                    // 201 — FIX: use explicit values
    StatusError                       // 202
)

// Better approach for HTTP-like enums:
const (
    StatusOK    HttpStatus = 200
    StatusNotFound      = 404
    StatusError         = 500
)

// iota is ideal for sequential enums:
type Priority int

const (
    PriorityLow Priority = iota  // 0
    PriorityMedium               // 1
    PriorityHigh                 // 2
    PriorityCritical             // 3
)
```

---

## Functions

### Multiple Return Values (No more `Pair<T,U>` or DTOs for returns)

```java
// Java: Need a wrapper class or custom type
public class Result {
    String data;
    Exception error;
}
public Result readFile(String path) { ... }
```

```go
// Go: Native multiple returns
func readFile(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(data), nil
}
```

### Named Return Values

```go
func divide(a, b float64) (result float64, err error) {
    if b == 0 {
        err = errors.New("division by zero")
        return  // naked return — uses named returns
    }
    result = a / b
    return
}
```

> ⚠️ **Best Practice**: Use named returns for documentation, but avoid naked returns in complex functions.

### Functions Are First-Class Citizens

```go
// Functions as values
type Transformer func(string) string

func process(data string, fn Transformer) string {
    return fn(data)
}

// Immediately Invoked Function Expression (IIFE)
result := func(x int) int {
    return x * 2
}(5)  // result = 10

// Closures (like Java lambdas, but simpler)
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c := counter()
fmt.Println(c())  // 1
fmt.Println(c())  // 2
```

### Variadic Functions (Like Java's `String...`)

```java
// Java
public int sum(int... numbers) { ... }
```

```go
// Go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

sum(1, 2, 3)       // 6
sum(1, 2, 3, 4, 5) // 15
```

### Deferred Calls (`defer`)

The `defer` keyword schedules a function call to run when the surrounding
function returns. It is commonly used for cleanup, such as closing a file or
releasing a resource, so cleanup still happens on early returns.

```go
func processFile(path string) error {
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    // Read or process file here.
    return nil
}
```

Deferred calls run in last-in, first-out (LIFO) order. The deferred function's
arguments are evaluated when the `defer` statement runs, while the call itself
waits until the surrounding function returns.

---

## Built-in Types

| Java Type | Go Type | Notes |
|-----------|---------|-------|
| `String` | `string` | UTF-8 by default, immutable |
| `int`/`Integer` | `int` | Platform-dependent (32 or 64 bit) |
| `long`/`Long` | `int64` | Explicit 64-bit |
| `double`/`Double` | `float64` | Default floating point |
| `float`/`Float` | `float32` | Use only when needed |
| `boolean`/`Boolean` | `bool` | `true` or `false` |
| `byte` | `byte` | Alias for `uint8` |
| `char` | `rune` | Alias for `int32` (Unicode code point) |

### Strings in Go

```go
// Strings are UTF-8 byte slices
s := "Hello, 世界"

// Length in bytes (NOT characters!)
fmt.Println(len(s))  // 13 (not 9!)

// Iterating by rune (character)
for i, r := range s {
    fmt.Printf("Index %d: %c\n", i, r)
}

// String to rune slice
runes := []rune(s)
fmt.Println(len(runes))  // 9 — correct character count
```

### Type Conversions (No Implicit Casting!)

```java
// Java: Implicit widening
int x = 42;
long y = x;  // OK
```

```go
// Go: EXPLICIT conversion required
var x int = 42
var y int64 = int64(x)  // Must be explicit!

// String conversions
s := strconv.Itoa(42)       // int → string
n, err := strconv.Atoi("42") // string → int
```

---

## Control Flow

### If Statements (No Parentheses!)

```java
// Java
if (x > 10) {
    doSomething();
}
```

```go
// Go
if x > 10 {
    doSomething()
}

// Go special: if with init statement
if err := doSomething(); err != nil {
    log.Fatal(err)
}
// err is scoped to this if/else block only
```

### Loops — Only `for`, No `while` or `do-while`

```java
// Java has for, while, do-while, for-each
for (int i = 0; i < 10; i++) { ... }
while (condition) { ... }
for (String item : list) { ... }
```

```go
// Go: Only 'for' — but it does everything

// Classic for
for i := 0; i < 10; i++ { ... }

// While-like
for condition { ... }

// Infinite loop
for { ... }

// For-each (range)
for index, value := range slice { ... }
for key, value := range myMap { ... }

// Ignore index or value with _
for _, value := range slice { ... }
```

### Switch — No `break` Needed, No Fall-through by Default

```java
// Java
switch (day) {
    case "Monday":   // Falls through!
    case "Tuesday":
        work();
        break;       // Must break explicitly
    default:
        rest();
}
```

```go
// Go: No fall-through by default
switch day {
case "Monday", "Tuesday":
    work()  // No break needed
default:
    rest()
}

// Type switch (no Java equivalent)
switch v := value.(type) {
case string:
    fmt.Println("string:", v)
case int:
    fmt.Println("int:", v)
default:
    fmt.Println("unknown type")
}
```

---

## Packages and Imports

### Java Package vs Go Package

```java
// Java: package + class + visibility modifiers
package com.example.service;

public class UserService {
    private String secret;    // class-private
    protected String internal; // package + subclass
    public String name;       // everyone
}
```

```go
// Go: package + file + capitalization for visibility
package service

type UserService struct {
    secret   string  // unexported (private to package)
    Internal string  // exported (public)
    Name     string  // exported (public)
}
```

> 🔑 **Key Rule**: In Go, visibility is **package-level**, not class-level. Capitalized = exported (public). Lowercase = unexported (package-private).

### Import Aliases and Blank Imports

```go
import (
    "fmt"                              // standard library
    "net/http"                         // standard library
    
    "github.com/gofiber/fiber/v2"     // third-party
    
    myjson "encoding/json"             // alias
    _ "github.com/lib/pq"             // blank import (side effects only)
)
```

---

## Key Takeaways for Java Developers

1. **`:=` is your friend** — Use short declaration inside functions
2. **No `null` for basic types** — Everything has a zero value
3. **`_` is the "ignore" operator** — Use it to discard unwanted return values
4. **No implicit type casting** — Always be explicit
5. **Only `for` loops** — No `while`, no `do-while`
6. **Capitalization = visibility** — No `public`/`private` keywords
7. **Functions are first-class** — Like Java lambdas but more powerful
8. **Multiple returns** — Replace DTO wrapper classes

---

**Practice**: [Module 01 Exercises →](../exercises/module-01-basics/)

Next: [Chapter 02 — Structs & Methods →](02-structs-and-methods.md)