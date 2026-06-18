# Chapter 04: Arrays, Slices & Maps

> No ArrayList, no HashMap. Go's collections are simpler and more efficient.

---

## Arrays — Fixed Size (Rarely Used Directly)

```java
// Java
int[] numbers = new int[5];
String[] names = {"Alice", "Bob"};
```

```go
// Go: Arrays have fixed size — part of the type!
var numbers [5]int                    // [0,0,0,0,0]
names := [2]string{"Alice", "Bob"}   // Size is fixed
fruits := [...]string{"Apple", "Banana"} // Compiler counts: [2]string

// [5]int and [10]int are DIFFERENT TYPES
var a [5]int
var b [10]int
// a = b  // COMPILE ERROR: different types!
```

> 🔑 **Key Insight**: Arrays are rarely used directly in Go. You'll almost always use **slices** instead.

---

## Slices — Go's "ArrayList"

### Memory Layout

From the ultimate-go resource, a slice is a 3-word data structure:

```
Slice Header (on stack)
  --------
 | ptr    | --> Backing Array (on heap)
  --------
 | len    |    3 (visible elements)
  --------
 | cap    |    5 (total capacity)
  --------
```

### Creating Slices

```go
// From the ultimate-go resource on slices:
// Slice literal
fruits := []string{"Apple", "Banana", "Cherry"}  // len=3, cap=3

// Using make
nums := make([]int, 5)        // len=5, cap=5, all zeros
nums := make([]int, 0, 10)    // len=0, cap=10 (pre-allocated)

// Slice from array
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2, 3, 4] — shares backing array!
```

### Append and Capacity

```go
slice := make([]int, 0, 3)

slice = append(slice, 1)  // len=1, cap=3
slice = append(slice, 2)  // len=2, cap=3
slice = append(slice, 3)  // len=3, cap=3

// Capacity exceeded — new backing array allocated (usually 2x)
slice = append(slice, 4)  // len=4, cap=6 (new array!)

// ALWAYS assign result of append:
slice = append(slice, 99)
```

> 🎯 **High-Traffic Insight**: Pre-allocate slices with `make([]T, 0, estimatedSize)` to avoid repeated allocations. Each `append` that exceeds capacity causes a new allocation and copy.

### Slicing Gotchas

```go
// WARNING: Slices share backing arrays!
original := []int{1, 2, 3, 4, 5}
sub := original[1:3]  // [2, 3]

sub[0] = 99
fmt.Println(original)  // [1, 99, 3, 4, 5] — original modified!

// Fix: Use copy or full slice expression
sub2 := make([]int, 2)
copy(sub2, original[1:3])

// Or use full slice expression (limits capacity)
sub3 := original[1:3:3]  // len=2, cap=2 — safe to append
```

### Copy

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)
n := copy(dst, src)  // Copies min(len(dst), len(src)) elements
// dst = [1, 2, 3], n = 3
```

---

## Maps — Go's "HashMap"

### Creating Maps

```java
// Java
Map<String, Integer> ages = new HashMap<>();
ages.put("Alice", 30);
ages.get("Alice");  // 30
ages.get("Bob");    // null
```

```go
// Go
ages := make(map[string]int)
ages["Alice"] = 30
ages["Alice"]       // 30
ages["Bob"]         // 0 (zero value for int — same as missing!)

// Check if key exists
age, ok := ages["Bob"]
if !ok {
    // Key doesn't exist
}

// Map literal
colors := map[string]string{
    "red":   "#FF0000",
    "green": "#00FF00",
}
```

### Map Operations

```go
// CRUD
ages["Charlie"] = 25     // Create/Update
age := ages["Alice"]     // Read
delete(ages, "Charlie")  // Delete

// Iteration (RANDOM order — important!)
for key, value := range ages {
    fmt.Println(key, value)
}

// Length
len(ages)  // Number of key-value pairs
```

### Map Iteration Order Is Random

> ⚠️ Unlike Java's LinkedHashMap, Go maps have **no guaranteed iteration order**. The Go runtime intentionally randomizes iteration to prevent dependency on order.

```go
// If you need sorted keys:
keys := make([]string, 0, len(ages))
for k := range ages {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, ages[k])
}
```

### Map Gotcha: Values Are Not Addressable

```go
type Point struct{ X, Y int }
m := map[string]Point{"origin": {0, 0}}

// m["origin"].X = 5  // COMPILE ERROR: cannot assign to struct field in map

// Fix: Store pointer, or reassign whole value
p := m["origin"]
p.X = 5
m["origin"] = p

// Or use pointer values:
m2 := map[string]*Point{}
m2["origin"] = &Point{0, 0}
m2["origin"].X = 5  // OK
```

---

## Java Collections → Go Cheat Sheet

| Java | Go | Notes |
|------|----|-------|
| `ArrayList<T>` | `[]T` (slice) | Dynamic, use `append()` |
| `LinkedList<T>` | `container/list` | Rarely used |
| `HashMap<K,V>` | `map[K]V` | Built-in, not thread-safe |
| `TreeMap<K,V>` | No built-in | Use sorted slice or 3rd party |
| `HashSet<T>` | `map[T]struct{}` | Use empty struct for value |
| `LinkedHashMap` | No built-in | Maintain separate slice of keys |
| `Arrays.sort()` | `sort.Slice()` | In-place sorting |
| `Stream.map()` | `for range` loop | No stream API in Go |
| `Stream.filter()` | `for range` with `if` | No stream API in Go |

---

## Key Takeaways

1. **Use slices, not arrays** — Slices are dynamic and flexible
2. **Pre-allocate with `make`** — Critical for high-traffic performance
3. **Slices share backing arrays** — Be careful with sub-slices
4. **Maps have random iteration order** — Never depend on order
5. **Check `ok` for map lookups** — Zero value and missing key return the same thing
6. **Use `map[T]struct{}` for sets** — Zero memory per value

---

**Practice**: [Module 02 — Collections Exercises →](../exercises/module-02-collections/)

Next: [Chapter 05 — Error Handling →](05-error-handling.md)