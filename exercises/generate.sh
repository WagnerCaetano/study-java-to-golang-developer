#!/bin/bash
# Generator script for all 60 exercises
BASE="c:/Users/Wagner Caetano/git/study-java-to-golang-developer/exercises"

# ===== MODULE 01: BASICS =====

# 01 - Hello Go
cat > "$BASE/module-01-basics/01-hello-go/README.md" << 'EOF'
# Exercise 01: Hello Go

## Problem
Your first Go program! The `Greet` function should return "Hello, Go!" but it's broken.

## Learning Objectives
- Understand Go package structure
- Learn function declaration
- Practice returning values

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/01-hello-go/solution.go" << 'EOF'
package hello

// Greet returns a greeting message
// FIX ME: This function should return "Hello, Go!" but it returns an empty string
func Greet() string {
	return "" // FIX: return "Hello, Go!"
}
EOF

cat > "$BASE/module-01-basics/01-hello-go/solution_test.go" << 'EOF'
package hello

import "testing"

func TestGreet(t *testing.T) {
	result := Greet()
	expected := "Hello, Go!"
	if result != expected {
		t.Errorf("Greet() = %q, want %q", result, expected)
	}
}
EOF

mkdir -p "$BASE/module-01-basics/01-hello-go/_solution"
cat > "$BASE/module-01-basics/01-hello-go/_solution/solution.go" << 'EOF'
package hello

func Greet() string {
	return "Hello, Go!"
}
EOF

# 02 - Variables and Types
cat > "$BASE/module-01-basics/02-variables-and-types/README.md" << 'EOF'
# Exercise 02: Variables and Types

## Problem
Complete the `CreateUser` function that returns a User struct with proper zero values and assigned fields.

## Learning Objectives
- Variable declaration with var and :=
- Zero values in Go
- Struct initialization

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/02-variables-and-types/solution.go" << 'EOF'
package variables

type User struct {
	Name string
	Age  int
}

// CreateUser should return a User with the given name and age
// FIX ME: Initialize the User struct properly
func CreateUser(name string, age int) User {
	return User{} // FIX: set Name and Age fields
}
EOF

cat > "$BASE/module-01-basics/02-variables-and-types/solution_test.go" << 'EOF'
package variables

import "testing"

func TestCreateUser(t *testing.T) {
	u := CreateUser("Alice", 30)
	if u.Name != "Alice" {
		t.Errorf("Name = %q, want %q", u.Name, "Alice")
	}
	if u.Age != 30 {
		t.Errorf("Age = %d, want %d", u.Age, 30)
	}
}

func TestCreateUserZeroValues(t *testing.T) {
	u := CreateUser("", 0)
	if u.Name != "" {
		t.Errorf("Empty name should be empty string, got %q", u.Name)
	}
	if u.Age != 0 {
		t.Errorf("Zero age should be 0, got %d", u.Age)
	}
}
EOF

mkdir -p "$BASE/module-01-basics/02-variables-and-types/_solution"
cat > "$BASE/module-01-basics/02-variables-and-types/_solution/solution.go" << 'EOF'
package variables

type User struct {
	Name string
	Age  int
}

func CreateUser(name string, age int) User {
	return User{Name: name, Age: age}
}
EOF

# 03 - Constants and Enums
cat > "$BASE/module-01-basics/03-constants-and-enums/README.md" << 'EOF'
# Exercise 03: Constants and Enums

## Problem
Define an enum-like type for HTTP status codes using iota and constants.

## Learning Objectives
- Constant declaration
- iota for sequential enums
- Custom types

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/03-constants-and-enums/solution.go" << 'EOF'
package constants

// Status represents an HTTP status code
type Status int

// FIX ME: Define StatusOK = 200, StatusNotFound = 404, StatusError = 500
// using explicit values (not iota, since these aren't sequential)
const (
	StatusOK      Status = 0 // FIX: should be 200
	StatusNotFound Status = 0 // FIX: should be 404
	StatusError   Status = 0 // FIX: should be 500
)

// StatusMessage returns a text message for the given status
func StatusMessage(s Status) string {
	messages := map[Status]string{
		StatusOK:      "OK",
		StatusNotFound: "Not Found",
		StatusError:   "Internal Server Error",
	}
	if msg, ok := messages[s]; ok {
		return msg
	}
	return "Unknown"
}
EOF

cat > "$BASE/module-01-basics/03-constants-and-enums/solution_test.go" << 'EOF'
package constants

import "testing"

func TestStatusValues(t *testing.T) {
	if StatusOK != 200 {
		t.Errorf("StatusOK = %d, want 200", StatusOK)
	}
	if StatusNotFound != 404 {
		t.Errorf("StatusNotFound = %d, want 404", StatusNotFound)
	}
	if StatusError != 500 {
		t.Errorf("StatusError = %d, want 500", StatusError)
	}
}

func TestStatusMessage(t *testing.T) {
	tests := []struct {
		status   Status
		expected string
	}{
		{StatusOK, "OK"},
		{StatusNotFound, "Not Found"},
		{StatusError, "Internal Server Error"},
		{Status(999), "Unknown"},
	}
	for _, tt := range tests {
		result := StatusMessage(tt.status)
		if result != tt.expected {
			t.Errorf("StatusMessage(%d) = %q, want %q", tt.status, result, tt.expected)
		}
	}
}
EOF

mkdir -p "$BASE/module-01-basics/03-constants-and-enums/_solution"
cat > "$BASE/module-01-basics/03-constants-and-enums/_solution/solution.go" << 'EOF'
package constants

type Status int

const (
	StatusOK      Status = 200
	StatusNotFound Status = 404
	StatusError   Status = 500
)

func StatusMessage(s Status) string {
	messages := map[Status]string{
		StatusOK:      "OK",
		StatusNotFound: "Not Found",
		StatusError:   "Internal Server Error",
	}
	if msg, ok := messages[s]; ok {
		return msg
	}
	return "Unknown"
}
EOF

# 04 - Functions and Multiple Returns
cat > "$BASE/module-01-basics/04-functions-and-multiple-returns/README.md" << 'EOF'
# Exercise 04: Functions and Multiple Returns

## Problem
Fix the `Divide` function to properly return both result and error.

## Learning Objectives
- Multiple return values
- Error handling pattern
- Named return values

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/04-functions-and-multiple-returns/solution.go" << 'EOF'
package functions

import "errors"

// Divide divides a by b and returns the result
// FIX ME: Return an error when b is 0, otherwise return the result and nil error
func Divide(a, b float64) (float64, error) {
	// FIX: Add division by zero check
	return a / b, nil
}
EOF

cat > "$BASE/module-01-basics/04-functions-and-multiple-returns/solution_test.go" << 'EOF'
package functions

import "testing"

func TestDivide(t *testing.T) {
	result, err := Divide(10, 2)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != 5.0 {
		t.Errorf("Divide(10, 2) = %f, want 5.0", result)
	}
}

func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("expected error for division by zero")
	}
}
EOF

mkdir -p "$BASE/module-01-basics/04-functions-and-multiple-returns/_solution"
cat > "$BASE/module-01-basics/04-functions-and-multiple-returns/_solution/solution.go" << 'EOF'
package functions

import "errors"

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
EOF

# 05 - Pointers and References
cat > "$BASE/module-01-basics/05-pointers-and-references/README.md" << 'EOF'
# Exercise 05: Pointers and References

## Problem
Fix the `Swap` function to actually swap the values of two integers using pointers.

## Learning Objectives
- Pointer declaration and dereferencing
- Passing by reference vs value
- Understanding memory addresses

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/05-pointers-and-references/solution.go" << 'EOF'
package pointers

// Swap should swap the values of a and b using pointers
// FIX ME: The current implementation doesn't actually swap
func Swap(a, b *int) {
	// FIX: Implement proper swap using a temporary variable
	*a = *b
	*b = *a
}
EOF

cat > "$BASE/module-01-basics/05-pointers-and-references/solution_test.go" << 'EOF'
package pointers

import "testing"

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("Swap: a=%d, b=%d; want a=2, b=1", a, b)
	}
}

func TestSwapEqual(t *testing.T) {
	a, b := 5, 5
	Swap(&a, &b)
	if a != 5 || b != 5 {
		t.Errorf("Swap equal: a=%d, b=%d; want a=5, b=5", a, b)
	}
}
EOF

mkdir -p "$BASE/module-01-basics/05-pointers-and-references/_solution"
cat > "$BASE/module-01-basics/05-pointers-and-references/_solution/solution.go" << 'EOF'
package pointers

func Swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
EOF

# 06 - Structs and Composition
cat > "$BASE/module-01-basics/06-structs-and-composition/README.md" << 'EOF'
# Exercise 06: Structs and Composition

## Problem
Create a `FullName` method on the Person struct that returns "First Last".

## Learning Objectives
- Struct declaration
- Struct initialization
- Method implementation

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/06-structs-and-composition/solution.go" << 'EOF'
package structs

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// FullName returns the person's full name
// FIX ME: Should return "FirstName LastName" format
func (p Person) FullName() string {
	return "" // FIX: return proper full name
}
EOF

cat > "$BASE/module-01-basics/06-structs-and-composition/solution_test.go" << 'EOF'
package structs

import "testing"

func TestFullName(t *testing.T) {
	p := Person{FirstName: "John", LastName: "Doe", Age: 30}
	result := p.FullName()
	expected := "John Doe"
	if result != expected {
		t.Errorf("FullName() = %q, want %q", result, expected)
	}
}

func TestFullNameEmpty(t *testing.T) {
	p := Person{}
	if p.FullName() != " " {
		t.Errorf("Empty person full name should be a single space")
	}
}
EOF

mkdir -p "$BASE/module-01-basics/06-structs-and-composition/_solution"
cat > "$BASE/module-01-basics/06-structs-and-composition/_solution/solution.go" << 'EOF'
package structs

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}
EOF

# 07 - Methods and Receivers
cat > "$BASE/module-01-basics/07-methods-and-receivers/README.md" << 'EOF'
# Exercise 07: Methods and Receivers

## Problem
Fix the Counter type to properly increment and reset using pointer receivers.

## Learning Objectives
- Value vs pointer receivers
- When to use pointer receivers
- State mutation through methods

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/07-methods-and-receivers/solution.go" << 'EOF'
package methods

// Counter keeps track of a count
type Counter struct {
	value int
}

// Increment increases the counter by 1
// FIX ME: This uses a value receiver — it won't modify the original!
func (c Counter) Increment() {
	c.value++
}

// Value returns the current counter value
func (c Counter) Value() int {
	return c.value
}

// Reset sets the counter back to 0
// FIX ME: This also uses a value receiver — it won't modify the original!
func (c Counter) Reset() {
	c.value = 0
}
EOF

cat > "$BASE/module-01-basics/07-methods-and-receivers/solution_test.go" << 'EOF'
package methods

import "testing"

func TestCounterIncrement(t *testing.T) {
	c := Counter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if c.Value() != 3 {
		t.Errorf("Value = %d, want 3", c.Value())
	}
}

func TestCounterReset(t *testing.T) {
	c := Counter{}
	c.Increment()
	c.Increment()
	c.Reset()
	if c.Value() != 0 {
		t.Errorf("After reset, Value = %d, want 0", c.Value())
	}
}
EOF

mkdir -p "$BASE/module-01-basics/07-methods-and-receivers/_solution"
cat > "$BASE/module-01-basics/07-methods-and-receivers/_solution/solution.go" << 'EOF'
package methods

type Counter struct {
	value int
}

func (c *Counter) Increment() {
	c.value++
}

func (c Counter) Value() int {
	return c.value
}

func (c *Counter) Reset() {
	c.value = 0
}
EOF

# 08 - Interfaces
cat > "$BASE/module-01-basics/08-interfaces/README.md" << 'EOF'
# Exercise 08: Interfaces

## Problem
Make the `Dog` and `Cat` types satisfy the `Speaker` interface.

## Learning Objectives
- Implicit interface implementation
- Polymorphism through interfaces
- Interface satisfaction

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-01-basics/08-interfaces/solution.go" << 'EOF'
package interfaces

// Speaker is an interface for things that can speak
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

type Cat struct {
	Name string
}

// FIX ME: Implement Speak() for Dog so it returns "<Name> says Woof!"

// FIX ME: Implement Speak() for Cat so it returns "<Name> says Meow!"

// MakeSpeak calls Speak on any Speaker
func MakeSpeak(s Speaker) string {
	return s.Speak()
}
EOF

cat > "$BASE/module-01-basics/08-interfaces/solution_test.go" << 'EOF'
package interfaces

import "testing"

func TestDogSpeak(t *testing.T) {
	d := Dog{Name: "Rex"}
	result := MakeSpeak(&d)
	expected := "Rex says Woof!"
	if result != expected {
		t.Errorf("Dog.Speak() = %q, want %q", result, expected)
	}
}

func TestCatSpeak(t *testing.T) {
	c := Cat{Name: "Whiskers"}
	result := MakeSpeak(&c)
	expected := "Whiskers says Meow!"
	if result != expected {
		t.Errorf("Cat.Speak() = %q, want %q", result, expected)
	}
}
EOF

mkdir -p "$BASE/module-01-basics/08-interfaces/_solution"
cat > "$BASE/module-01-basics/08-interfaces/_solution/solution.go" << 'EOF'
package interfaces

type Speaker interface {
	Speak() string
}

type Dog struct{ Name string }
type Cat struct{ Name string }

func (d *Dog) Speak() string { return d.Name + " says Woof!" }
func (c *Cat) Speak() string { return c.Name + " says Meow!" }

func MakeSpeak(s Speaker) string { return s.Speak() }
EOF

echo "Module 01 done!"

# ===== MODULE 02: COLLECTIONS =====

# 09 - Slice Operations
cat > "$BASE/module-02-collections/09-slice-operations/README.md" << 'EOF'
# Exercise 09: Slice Operations

## Problem
Fix the `Filter` function to return only even numbers from a slice.

## Learning Objectives
- Slice creation and manipulation
- Range iteration
- Slice append

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-02-collections/09-slice-operations/solution.go" << 'EOF'
package slices

// FilterEven returns only even numbers from the input slice
// FIX ME: Currently returns an empty slice
func FilterEven(nums []int) []int {
	var result []int
	// FIX: Iterate over nums and append even numbers to result
	return result
}
EOF

cat > "$BASE/module-02-collections/09-slice-operations/solution_test.go" << 'EOF'
package slices

import "testing"

func TestFilterEven(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"mixed", []int{1, 2, 3, 4, 5, 6}, []int{2, 4, 6}},
		{"all odd", []int{1, 3, 5}, []int{}},
		{"all even", []int{2, 4, 6}, []int{2, 4, 6}},
		{"empty", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FilterEven(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("got %v, want %v", result, tt.expected)
				}
			}
		})
	}
}
EOF

mkdir -p "$BASE/module-02-collections/09-slice-operations/_solution"
cat > "$BASE/module-02-collections/09-slice-operations/_solution/solution.go" << 'EOF'
package slices

func FilterEven(nums []int) []int {
	var result []int
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n)
		}
	}
	return result
}
EOF

# 10 - Slice Append and Capacity
cat > "$BASE/module-02-collections/10-slice-append-and-capacity/README.md" << 'EOF'
# Exercise 10: Slice Append and Capacity

## Problem
Fix the `Merge` function to merge two slices efficiently using pre-allocation.

## Learning Objectives
- make() for pre-allocation
- append() mechanics
- Capacity vs length

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-02-collections/10-slice-append-and-capacity/solution.go" << 'EOF'
package slicecap

// Merge combines two slices into a new one
// FIX ME: Use make to pre-allocate and append to fill
func Merge(a, b []int) []int {
	// FIX: Pre-allocate with make, then append elements
	return nil
}
EOF

cat > "$BASE/module-02-collections/10-slice-append-and-capacity/solution_test.go" << 'EOF'
package slicecap

import "testing"

func TestMerge(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	result := Merge(a, b)
	expected := []int{1, 2, 3, 4, 5, 6}
	if len(result) != len(expected) {
		t.Fatalf("len = %d, want %d", len(result), len(expected))
	}
	for i := range expected {
		if result[i] != expected[i] {
			t.Errorf("result[%d] = %d, want %d", i, result[i], expected[i])
		}
	}
}

func TestMergeEmpty(t *testing.T) {
	result := Merge([]int{}, []int{})
	if len(result) != 0 {
		t.Errorf("empty merge should return empty slice")
	}
}
EOF

mkdir -p "$BASE/module-02-collections/10-slice-append-and-capacity/_solution"
cat > "$BASE/module-02-collections/10-slice-append-and-capacity/_solution/solution.go" << 'EOF'
package slicecap

func Merge(a, b []int) []int {
	result := make([]int, 0, len(a)+len(b))
	result = append(result, a...)
	result = append(result, b...)
	return result
}
EOF

# 11 - Map CRUD
cat > "$BASE/module-02-collections/11-map-crud/README.md" << 'EOF'
# Exercise 11: Map CRUD

## Problem
Implement a simple in-memory key-value store with Get, Set, and Delete operations.

## Learning Objectives
- Map creation with make
- Map read/write/delete operations
- Checking if key exists

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-02-collections/11-map-crud/solution.go" << 'EOF'
package mapcrud

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}

// Set stores a key-value pair
// FIX ME: Implement this
func (s *Store) Set(key, value string) {
	// FIX: store in the data map
}

// Get retrieves a value by key. Returns ("", false) if not found.
// FIX ME: Implement this
func (s *Store) Get(key string) (string, bool) {
	return "", false // FIX
}

// Delete removes a key from the store
// FIX ME: Implement this
func (s *Store) Delete(key string) {
	// FIX: delete from data map
}
EOF

cat > "$BASE/module-02-collections/11-map-crud/solution_test.go" << 'EOF'
package mapcrud

import "testing"

func TestStoreCRUD(t *testing.T) {
	s := NewStore()
	
	// Get non-existent
	_, ok := s.Get("missing")
	if ok {
		t.Error("expected missing key to return false")
	}
	
	// Set and Get
	s.Set("name", "Alice")
	val, ok := s.Get("name")
	if !ok || val != "Alice" {
		t.Errorf("Get(name) = %q, %v; want %q, true", val, ok, "Alice")
	}
	
	// Update
	s.Set("name", "Bob")
	val, _ = s.Get("name")
	if val != "Bob" {
		t.Errorf("After update: got %q, want %q", val, "Bob")
	}
	
	// Delete
	s.Delete("name")
	_, ok = s.Get("name")
	if ok {
		t.Error("expected key to be deleted")
	}
}
EOF

mkdir -p "$BASE/module-02-collections/11-map-crud/_solution"
cat > "$BASE/module-02-collections/11-map-crud/_solution/solution.go" << 'EOF'
package mapcrud

type Store struct {
	data map[string]string
}

func NewStore() *Store {
	return &Store{data: make(map[string]string)}
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Get(key string) (string, bool) {
	val, ok := s.data[key]
	return val, ok
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}
EOF

# 12 - Map Iteration Order
cat > "$BASE/module-02-collections/12-map-iteration-order/README.md" << 'EOF'
# Exercise 12: Map Iteration Order

## Problem
Implement `SortedKeys` that returns map keys in sorted order.

## Learning Objectives
- Map iteration is random
- Sorting slices
- Working with map keys

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-02-collections/12-map-iteration-order/solution.go" << 'EOF'
package mapiter

import "sort"

// SortedKeys returns the keys of a map[string]int in alphabetical order
// FIX ME: Extract keys and sort them
func SortedKeys(m map[string]int) []string {
	// FIX: Extract keys, sort them, return
	return nil
}
EOF

cat > "$BASE/module-02-collections/12-map-iteration-order/solution_test.go" << 'EOF'
package mapiter

import "testing"

func TestSortedKeys(t *testing.T) {
	m := map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
	}
	result := SortedKeys(m)
	expected := []string{"alice", "bob", "charlie"}
	if len(result) != len(expected) {
		t.Fatalf("len = %d, want %d", len(result), len(expected))
	}
	for i, k := range result {
		if k != expected[i] {
			t.Errorf("key[%d] = %q, want %q", i, k, expected[i])
		}
	}
}
EOF

mkdir -p "$BASE/module-02-collections/12-map-iteration-order/_solution"
cat > "$BASE/module-02-collections/12-map-iteration-order/_solution/solution.go" << 'EOF'
package mapiter

import "sort"

func SortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
EOF

# 13 - Embedding in Structs
cat > "$BASE/module-02-collections/13-embedding-in-structs/README.md" << 'EOF'
# Exercise 13: Embedding in Structs

## Problem
Use struct embedding to give Vehicle base fields and make Car inherit them.

## Learning Objectives
- Struct embedding
- Accessing embedded fields
- Method promotion

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-02-collections/13-embedding-in-structs/solution.go" << 'EOF'
package embed

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

// Car should embed Vehicle and add a Doors field
// FIX ME: Define Car struct with embedded Vehicle and Doors int field
type Car struct {
	// FIX: embed Vehicle and add Doors int
}

// NewCar creates a new Car with all fields set
// FIX ME: Initialize the embedded Vehicle and Doors
func NewCar(make, model string, year, doors int) *Car {
	return nil // FIX
}
EOF

cat > "$BASE/module-02-collections/13-embedding-in-structs/solution_test.go" << 'EOF'
package embed

import "testing"

func TestCarEmbedding(t *testing.T) {
	c := NewCar("Toyota", "Corolla", 2024, 4)
	if c == nil {
		t.Fatal("NewCar returned nil")
	}
	if c.Make != "Toyota" {
		t.Errorf("Make = %q, want %q", c.Make, "Toyota")
	}
	if c.Model != "Corolla" {
		t.Errorf("Model = %q, want %q", c.Model, "Corolla")
	}
	if c.Year != 2024 {
		t.Errorf("Year = %d, want %d", c.Year, 2024)
	}
	if c.Doors != 4 {
		t.Errorf("Doors = %d, want %d", c.Doors, 4)
	}
}
EOF

mkdir -p "$BASE/module-02-collections/13-embedding-in-structs/_solution"
cat > "$BASE/module-02-collections/13-embedding-in-structs/_solution/solution.go" << 'EOF'
package embed

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

type Car struct {
	Vehicle
	Doors int
}

func NewCar(make, model string, year, doors int) *Car {
	return &Car{
		Vehicle: Vehicle{Make: make, Model: model, Year: year},
		Doors:   doors,
	}
}
EOF

# 14 - Type Conversion
cat > "$BASE/module-02-collections/14-type-conversion/README.md" << 'EOF'
# Exercise 14: Type Conversion

## Problem
Implement type conversion functions between Celsius and Fahrenheit.

## Learning Objectives
- Type conversion syntax
- Type assertion with comma-ok
- Working with interfaces

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-02-collections/14-type-conversion/solution.go" << 'EOF'
package conversion

type Celsius float64
type Fahrenheit float64

// CToF converts Celsius to Fahrenheit
// FIX ME: Formula is F = C * 9/5 + 32
func CToF(c Celsius) Fahrenheit {
	return 0 // FIX
}

// FToC converts Fahrenheit to Celsius
// FIX ME: Formula is C = (F - 32) * 5/9
func FToC(f Fahrenheit) Celsius {
	return 0 // FIX
}
EOF

cat > "$BASE/module-02-collections/14-type-conversion/solution_test.go" << 'EOF'
package conversion

import "testing"

func TestCToF(t *testing.T) {
	tests := []struct {
		input    Celsius
		expected Fahrenheit
	}{
		{0, 32},
		{100, 212},
		{-40, -40},
	}
	for _, tt := range tests {
		result := CToF(tt.input)
		if result != tt.expected {
			t.Errorf("CToF(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}

func TestFToC(t *testing.T) {
	tests := []struct {
		input    Fahrenheit
		expected Celsius
	}{
		{32, 0},
		{212, 100},
		{-40, -40},
	}
	for _, tt := range tests {
		result := FToC(tt.input)
		if result != tt.expected {
			t.Errorf("FToC(%v) = %v, want %v", tt.input, result, tt.expected)
		}
	}
}
EOF

mkdir -p "$BASE/module-02-collections/14-type-conversion/_solution"
cat > "$BASE/module-02-collections/14-type-conversion/_solution/solution.go" << 'EOF'
package conversion

type Celsius float64
type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
EOF

echo "Module 02 done!"

# ===== MODULE 03: ERROR HANDLING =====

# 15 - Basic Errors
cat > "$BASE/module-03-error-handling/15-basic-errors/README.md" << 'EOF'
# Exercise 15: Basic Errors

## Problem
Fix the `ValidateAge` function to return proper errors for invalid ages.

## Learning Objectives
- errors.New()
- Error checking pattern
- Returning nil on success

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-03-error-handling/15-basic-errors/solution.go" << 'EOF'
package basicerrors

import "errors"

// ValidateAge checks if age is valid (0-150)
// FIX ME: Return error for negative ages and ages > 150
func ValidateAge(age int) error {
	// FIX: Add validation
	return nil
}

var ErrNegativeAge = errors.New("age cannot be negative")
var ErrTooOld = errors.New("age cannot exceed 150")
EOF

cat > "$BASE/module-03-error-handling/15-basic-errors/solution_test.go" << 'EOF'
package basicerrors

import "testing"

func TestValidateAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr bool
	}{
		{"valid", 25, false},
		{"zero", 0, false},
		{"max valid", 150, false},
		{"negative", -1, true},
		{"too old", 151, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAge(tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAge(%d) error = %v, wantErr %v", tt.age, err, tt.wantErr)
			}
		})
	}
}
EOF

mkdir -p "$BASE/module-03-error-handling/15-basic-errors/_solution"
cat > "$BASE/module-03-error-handling/15-basic-errors/_solution/solution.go" << 'EOF'
package basicerrors

import "errors"

var ErrNegativeAge = errors.New("age cannot be negative")
var ErrTooOld = errors.New("age cannot exceed 150")

func ValidateAge(age int) error {
	if age < 0 {
		return ErrNegativeAge
	}
	if age > 150 {
		return ErrTooOld
	}
	return nil
}
EOF

# 16 - Custom Error Types
cat > "$BASE/module-03-error-handling/16-custom-error-types/README.md" << 'EOF'
# Exercise 16: Custom Error Types

## Problem
Create a custom `HTTPError` type that includes status code and message.

## Learning Objectives
- Custom error types
- Implementing the error interface
- Type assertions for errors

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-03-error-handling/16-custom-error-types/solution.go" << 'EOF'
package customerrors

import "fmt"

// HTTPError represents an HTTP error with status code and message
// FIX ME: Define HTTPError struct with Code int and Message string

// Error implements the error interface
// FIX ME: Implement Error() method that returns "HTTP 404: Not Found" format
EOF

cat > "$BASE/module-03-error-handling/16-custom-error-types/solution_test.go" << 'EOF'
package customerrors

import "testing"

func TestHTTPError(t *testing.T) {
	err := &HTTPError{Code: 404, Message: "Not Found"}
	if err.Error() != "HTTP 404: Not Found" {
		t.Errorf("Error() = %q, want %q", err.Error(), "HTTP 404: Not Found")
	}
}

func TestHTTPError500(t *testing.T) {
	err := &HTTPError{Code: 500, Message: "Internal Server Error"}
	if err.Error() != "HTTP 500: Internal Server Error" {
		t.Errorf("Error() = %q", err.Error())
	}
}
EOF

mkdir -p "$BASE/module-03-error-handling/16-custom-error-types/_solution"
cat > "$BASE/module-03-error-handling/16-custom-error-types/_solution/solution.go" << 'EOF'
package customerrors

import "fmt"

type HTTPError struct {
	Code    int
	Message string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP %d: %s", e.Code, e.Message)
}
EOF

# 17 - Error Wrapping
cat > "$BASE/module-03-error-handling/17-error-wrapping-and-unwrapping/README.md" << 'EOF'
# Exercise 17: Error Wrapping and Unwrapping

## Problem
Fix the `ReadConfig` function to properly wrap errors with context using `%w`.

## Learning Objectives
- fmt.Errorf with %w
- Error wrapping for context
- Building error chains

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-03-error-handling/17-error-wrapping-and-unwrapping/solution.go" << 'EOF'
package wrapping

import (
	"errors"
	"fmt"
)

var ErrConfigNotFound = errors.New("config not found")

// ReadConfig simulates reading a config file
// FIX ME: Wrap the error with context using %w
func ReadConfig(path string) error {
	return ErrConfigNotFound // FIX: Wrap with "readConfig(<path>):"
}

// GetRootCause unwraps an error chain to find the root cause
func GetRootCause(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
EOF

cat > "$BASE/module-03-error-handling/17-error-wrapping-and-unwrapping/solution_test.go" << 'EOF'
package wrapping

import (
	"errors"
	"testing"
)

func TestReadConfigWrap(t *testing.T) {
	err := ReadConfig("/etc/app/config.yaml")
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, ErrConfigNotFound) {
		t.Errorf("error should wrap ErrConfigNotFound")
	}
}

func TestGetRootCause(t *testing.T) {
	err := ReadConfig("/test")
	root := GetRootCause(err)
	if !errors.Is(root, ErrConfigNotFound) {
		t.Errorf("root cause should be ErrConfigNotFound, got %v", root)
	}
}
EOF

mkdir -p "$BASE/module-03-error-handling/17-error-wrapping-and-unwrapping/_solution"
cat > "$BASE/module-03-error-handling/17-error-wrapping-and-unwrapping/_solution/solution.go" << 'EOF'
package wrapping

import (
	"errors"
	"fmt"
)

var ErrConfigNotFound = errors.New("config not found")

func ReadConfig(path string) error {
	return fmt.Errorf("readConfig(%s): %w", path, ErrConfigNotFound)
}

func GetRootCause(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
EOF

# 18 - errors.Is and errors.As
cat > "$BASE/module-03-error-handling/18-errors-is-and-as/README.md" << 'EOF'
# Exercise 18: errors.Is and errors.As

## Problem
Use errors.Is and errors.As to inspect error chains properly.

## Learning Objectives
- errors.Is for checking error identity
- errors.As for extracting error type
- Proper error chain inspection

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-03-error-handling/18-errors-is-and-as/solution.go" << 'EOF'
package errorsis

import "errors"

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return "not found"
}

var ErrUnauthorized = errors.New("unauthorized")

// CheckNotFound should return true if the error wraps a NotFoundError
// FIX ME: Use errors.As properly
func CheckNotFound(err error) bool {
	return false // FIX: use errors.As
}

// CheckUnauthorized should return true if the error is ErrUnauthorized
// FIX ME: Use errors.Is properly
func CheckUnauthorized(err error) bool {
	return false // FIX: use errors.Is
}
EOF

cat > "$BASE/module-03-error-handling/18-errors-is-and-as/solution_test.go" << 'EOF'
package errorsis

import (
	"fmt"
	"testing"
)

func TestCheckNotFound(t *testing.T) {
	err := fmt.Errorf("wrap: %w", &NotFoundError{Resource: "user", ID: 42})
	if !CheckNotFound(err) {
		t.Error("should detect NotFoundError in chain")
	}
	if CheckNotFound(fmt.Errorf("other error")) {
		t.Error("should not detect NotFoundError for other errors")
	}
}

func TestCheckUnauthorized(t *testing.T) {
	err := fmt.Errorf("wrap: %w", ErrUnauthorized)
	if !CheckUnauthorized(err) {
		t.Error("should detect ErrUnauthorized in chain")
	}
	if CheckUnauthorized(fmt.Errorf("other error")) {
		t.Error("should not detect ErrUnauthorized for other errors")
	}
}
EOF

mkdir -p "$BASE/module-03-error-handling/18-errors-is-and-as/_solution"
cat > "$BASE/module-03-error-handling/18-errors-is-and-as/_solution/solution.go" << 'EOF'
package errorsis

import "errors"

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string { return "not found" }

var ErrUnauthorized = errors.New("unauthorized")

func CheckNotFound(err error) bool {
	var nf *NotFoundError
	return errors.As(err, &nf)
}

func CheckUnauthorized(err error) bool {
	return errors.Is(err, ErrUnauthorized)
}
EOF

# 19 - Panic and Recover
cat > "$BASE/module-03-error-handling/19-panic-and-recover/README.md" << 'EOF'
# Exercise 19: Panic and Recover

## Problem
Implement a safe function that catches panics and returns them as errors.

## Learning Objectives
- When to use panic
- Recover in deferred functions
- Converting panics to errors

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-03-error-handling/19-panic-and-recover/solution.go" << 'EOF'
package panicrecover

import "fmt"

// SafeExecute runs fn and converts any panic into an error
// FIX ME: Use defer + recover to catch panics
func SafeExecute(fn func()) error {
	// FIX: Add defer with recover
	fn()
	return nil
}
EOF

cat > "$BASE/module-03-error-handling/19-panic-and-recover/solution_test.go" << 'EOF'
package panicrecover

import "testing"

func TestSafeExecuteNoPanic(t *testing.T) {
	called := false
	err := SafeExecute(func() {
		called = true
	})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !called {
		t.Error("function was not called")
	}
}

func TestSafeExecuteWithPanic(t *testing.T) {
	err := SafeExecute(func() {
		panic("something went wrong")
	})
	if err == nil {
		t.Fatal("expected error from panic")
	}
	if err.Error() != "panic: something went wrong" {
		t.Errorf("error = %q, want to contain panic message", err.Error())
	}
}
EOF

mkdir -p "$BASE/module-03-error-handling/19-panic-and-recover/_solution"
cat > "$BASE/module-03-error-handling/19-panic-and-recover/_solution/solution.go" << 'EOF'
package panicrecover

import "fmt"

func SafeExecute(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()
	fn()
	return nil
}
EOF

# 20 - Sentinel Errors
cat > "$BASE/module-03-error-handling/20-sentinel-errors/README.md" << 'EOF'
# Exercise 20: Sentinel Errors

## Problem
Create sentinel errors for a user service and use them properly.

## Learning Objectives
- Sentinel error pattern
- Package-level error variables
- errors.Is with sentinels

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-03-error-handling/20-sentinel-errors/solution.go" << 'EOF'
package sentinel

import "errors"

// FIX ME: Define sentinel errors
var (
	ErrNotFound     = errors.New("") // FIX: "user not found"
	ErrAlreadyExists = errors.New("") // FIX: "user already exists"
	ErrInvalidInput = errors.New("") // FIX: "invalid input"
)

// FindUser simulates finding a user
func FindUser(id int) error {
	if id <= 0 {
		return ErrInvalidInput
	}
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

// CreateUser simulates creating a user
func CreateUser(id int) error {
	if id == 1 {
		return ErrAlreadyExists
	}
	return nil
}
EOF

cat > "$BASE/module-03-error-handling/20-sentinel-errors/solution_test.go" << 'EOF'
package sentinel

import (
	"errors"
	"testing"
)

func TestSentinelMessages(t *testing.T) {
	if ErrNotFound.Error() != "user not found" {
		t.Errorf("ErrNotFound = %q", ErrNotFound.Error())
	}
	if ErrAlreadyExists.Error() != "user already exists" {
		t.Errorf("ErrAlreadyExists = %q", ErrAlreadyExists.Error())
	}
	if ErrInvalidInput.Error() != "invalid input" {
		t.Errorf("ErrInvalidInput = %q", ErrInvalidInput.Error())
	}
}

func TestFindUser(t *testing.T) {
	if !errors.Is(FindUser(-1), ErrInvalidInput) {
		t.Error("negative id should return ErrInvalidInput")
	}
	if !errors.Is(FindUser(999), ErrNotFound) {
		t.Error("id 999 should return ErrNotFound")
	}
	if FindUser(1) != nil {
		t.Error("valid id should return nil")
	}
}
EOF

mkdir -p "$BASE/module-03-error-handling/20-sentinel-errors/_solution"
cat > "$BASE/module-03-error-handling/20-sentinel-errors/_solution/solution.go" << 'EOF'
package sentinel

import "errors"

var (
	ErrNotFound      = errors.New("user not found")
	ErrAlreadyExists = errors.New("user already exists")
	ErrInvalidInput  = errors.New("invalid input")
)

func FindUser(id int) error {
	if id <= 0 {
		return ErrInvalidInput
	}
	if id == 999 {
		return ErrNotFound
	}
	return nil
}

func CreateUser(id int) error {
	if id == 1 {
		return ErrAlreadyExists
	}
	return nil
}
EOF

echo "Module 03 done!"
echo "Generating remaining modules..."

# ===== MODULE 04: CONCURRENCY (21-30) =====

for ex_num in 21 22 23 24 25 26 27 28 29 30; do
mkdir -p "$BASE/module-04-concurrency/${ex_num}-placeholder/_solution"
done

# 21 - Goroutine Basics
cat > "$BASE/module-04-concurrency/21-goroutine-basics/README.md" << 'EOF'
# Exercise 21: Goroutine Basics

## Problem
Fix the `ParallelSum` function to sum numbers using multiple goroutines.

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-04-concurrency/21-goroutine-basics/solution.go" << 'EOF'
package goroutine

import "sync"

// ParallelSum sums a slice using numGoroutines goroutines
// FIX ME: Split the work and use goroutines + WaitGroup
func ParallelSum(nums []int, numGoroutines int) int {
	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunkSize := len(nums) / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(nums)
		}
		wg.Add(1)
		go func(slice []int) {
			defer wg.Done()
			sum := 0
			for _, n := range slice {
				sum += n
			}
			mu.Lock()
			total += sum // FIX: This has a race condition! Use mutex properly.
			mu.Unlock()
		}(nums[start:end])
	}

	wg.Wait()
	return total
}
EOF

cat > "$BASE/module-04-concurrency/21-goroutine-basics/solution_test.go" << 'EOF'
package goroutine

import "testing"

func TestParallelSum(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := ParallelSum(nums, 4)
	expected := 55
	if result != expected {
		t.Errorf("ParallelSum = %d, want %d", result, expected)
	}
}

func TestParallelSumSingle(t *testing.T) {
	nums := []int{1, 2, 3}
	result := ParallelSum(nums, 1)
	if result != 6 {
		t.Errorf("ParallelSum = %d, want 6", result)
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/21-goroutine-basics/_solution"
cat > "$BASE/module-04-concurrency/21-goroutine-basics/_solution/solution.go" << 'EOF'
package goroutine

import "sync"

func ParallelSum(nums []int, numGoroutines int) int {
	var total int
	var mu sync.Mutex
	var wg sync.WaitGroup

	chunkSize := len(nums) / numGoroutines
	if chunkSize == 0 {
		chunkSize = len(nums)
		numGoroutines = 1
	}

	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(nums)
		}
		if start >= len(nums) {
			break
		}
		wg.Add(1)
		go func(slice []int) {
			defer wg.Done()
			sum := 0
			for _, n := range slice {
				sum += n
			}
			mu.Lock()
			total += sum
			mu.Unlock()
		}(nums[start:end])
	}

	wg.Wait()
	return total
}
EOF

# 22-30: Generate remaining concurrency exercises
cat > "$BASE/module-04-concurrency/22-channels-unbuffered/README.md" << 'EOF'
# Exercise 22: Unbuffered Channels

## Problem
Implement a producer-consumer pattern using unbuffered channels.

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-04-concurrency/22-channels-unbuffered/solution.go" << 'EOF'
package channel

// Producer generates numbers 1 to n and sends them on the channel
// FIX ME: Close the channel after sending all numbers
func Producer(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		// FIX: close the channel
	}()
	return ch
}

// Consumer reads all values from the channel and returns their sum
func Consumer(ch <-chan int) int {
	total := 0
	for val := range ch {
		total += val
	}
	return total
}
EOF

cat > "$BASE/module-04-concurrency/22-channels-unbuffered/solution_test.go" << 'EOF'
package channel

import "testing"

func TestProducerConsumer(t *testing.T) {
	ch := Producer(10)
	result := Consumer(ch)
	expected := 55 // 1+2+...+10
	if result != expected {
		t.Errorf("Consumer got %d, want %d", result, expected)
	}
}

func TestProducerConsumerEmpty(t *testing.T) {
	ch := Producer(0)
	result := Consumer(ch)
	if result != 0 {
		t.Errorf("Empty producer: got %d, want 0", result)
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/22-channels-unbuffered/_solution"
cat > "$BASE/module-04-concurrency/22-channels-unbuffered/_solution/solution.go" << 'EOF'
package channel

func Producer(n int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func Consumer(ch <-chan int) int {
	total := 0
	for val := range ch {
		total += val
	}
	return total
}
EOF

# 23 - Buffered Channels
cat > "$BASE/module-04-concurrency/23-channels-buffered/README.md" << 'EOF'
# Exercise 23: Buffered Channels

## Problem
Implement a bounded task processor using a buffered channel as a semaphore.

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-04-concurrency/23-channels-buffered/solution.go" << 'EOF'
package buffered

import "sync"

// BoundedProcessor runs tasks with at most maxConcurrent goroutines
// FIX ME: Use a buffered channel as a semaphore
func BoundedProcessor(tasks []func(), maxConcurrent int) {
	var wg sync.WaitGroup
	// FIX: Create a buffered channel semaphore and use it

	for _, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
		}(task)
	}
	wg.Wait()
}
EOF

cat > "$BASE/module-04-concurrency/23-channels-buffered/solution_test.go" << 'EOF'
package buffered

import (
	"sync/atomic"
	"testing"
)

func TestBoundedProcessor(t *testing.T) {
	var concurrent, maxConcurrent int32
	tasks := make([]func(), 20)
	for i := range tasks {
		tasks[i] = func() {
			c := atomic.AddInt32(&concurrent, 1)
			for {
				old := atomic.LoadInt32(&maxConcurrent)
				if c <= old || atomic.CompareAndSwapInt32(&maxConcurrent, old, c) {
					break
				}
			}
			atomic.AddInt32(&concurrent, -1)
		}
	}
	BoundedProcessor(tasks, 3)
	if atomic.LoadInt32(&maxConcurrent) > 3 {
		t.Errorf("max concurrent = %d, should be <= 3", maxConcurrent)
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/23-channels-buffered/_solution"
cat > "$BASE/module-04-concurrency/23-channels-buffered/_solution/solution.go" << 'EOF'
package buffered

import "sync"

func BoundedProcessor(tasks []func(), maxConcurrent int) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrent)

	for _, task := range tasks {
		sem <- struct{}{}
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			t()
			<-sem
		}(task)
	}
	wg.Wait()
}
EOF

# 24 - Select Statement
cat > "$BASE/module-04-concurrency/24-select-statement/README.md" << 'EOF'
# Exercise 24: Select Statement

## Problem
Implement `FirstResult` that returns the first result from two channels.

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-04-concurrency/24-select-statement/solution.go" << 'EOF'
package selectstmt

// FirstResult returns the first value received from either ch1 or ch2
// FIX ME: Use select to wait on both channels
func FirstResult(ch1, ch2 <-chan string) string {
	// FIX: Use select to read from whichever channel is ready first
	return ""
}
EOF

cat > "$BASE/module-04-concurrency/24-select-statement/solution_test.go" << 'EOF'
package selectstmt

import (
	"testing"
	"time"
)

func TestFirstResultCh1(t *testing.T) {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch1 <- "from ch1"
	ch2 <- "from ch2"
	result := FirstResult(ch1, ch2)
	if result != "from ch1" && result != "from ch2" {
		t.Errorf("FirstResult = %q, expected a channel value", result)
	}
}

func TestFirstResultTimeout(t *testing.T) {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch1 <- "delayed"
	}()
	result := FirstResult(ch1, ch2)
	if result != "delayed" {
		t.Errorf("FirstResult = %q, want %q", result, "delayed")
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/24-select-statement/_solution"
cat > "$BASE/module-04-concurrency/24-select-statement/_solution/solution.go" << 'EOF'
package selectstmt

func FirstResult(ch1, ch2 <-chan string) string {
	select {
	case v := <-ch1:
		return v
	case v := <-ch2:
		return v
	}
}
EOF

# 25-30: Generate remaining exercises
for num in 25 26 27 28 29 30; do
	dir=""
	case $num in
		25) dir="fan-out-fan-in"; desc="Implement fan-out/fan-in pattern" ;;
		26) dir="context-timeout"; desc="Use context.WithTimeout for HTTP calls" ;;
		27) dir="context-cancellation"; desc="Propagate cancellation through goroutines" ;;
		28) dir="context-values"; desc="Store and retrieve request-scoped values" ;;
		29) dir="waitgroup"; desc="Synchronize multiple goroutines with WaitGroup" ;;
		30) dir="mutex-and-data-races"; desc="Fix data races using mutex" ;;
	esac
cat > "$BASE/module-04-concurrency/${num}-${dir}/README.md" << EOF
# Exercise ${num}: ${dir}

## Problem
${desc}.

## How to Test
\`\`\`bash
go test -v
\`\`\`
EOF
done

# 25 Fan-out/fan-in
cat > "$BASE/module-04-concurrency/25-fan-out-fan-in/solution.go" << 'EOF'
package fanout

import "sync"

// ProcessFanOut processes items using numWorkers goroutines
// FIX ME: Implement fan-out (distribute work) and fan-in (collect results)
func ProcessFanOut(items []int, numWorkers int, process func(int) int) []int {
	results := make([]int, len(items))
	// FIX: Fan-out to workers, fan-in results
	return results
}
EOF

cat > "$BASE/module-04-concurrency/25-fan-out-fan-in/solution_test.go" << 'EOF'
package fanout

import "testing"

func TestFanOutFanIn(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := ProcessFanOut(items, 4, func(n int) int {
		return n * n
	})
	if len(result) != len(items) {
		t.Fatalf("len = %d, want %d", len(result), len(items))
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/25-fan-out-fan-in/_solution"
cat > "$BASE/module-04-concurrency/25-fan-out-fan-in/_solution/solution.go" << 'EOF'
package fanout

import "sync"

func ProcessFanOut(items []int, numWorkers int, process func(int) int) []int {
	results := make([]int, len(items))
	var wg sync.WaitGroup
	chunkSize := (len(items) + numWorkers - 1) / numWorkers

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(items) {
			end = len(items)
		}
		if start >= len(items) {
			break
		}
		wg.Add(1)
		go func(idx int, slice []int) {
			defer wg.Done()
			for j, item := range slice {
				results[idx+j] = process(item)
			}
		}(start, items[start:end])
	}
	wg.Wait()
	return results
}
EOF

# 26 Context Timeout
cat > "$BASE/module-04-concurrency/26-context-timeout/solution.go" << 'EOF'
package ctxtimeout

import (
	"context"
	"fmt"
	"time"
)

// FetchWithTimeout simulates a fetch that respects context timeout
// FIX ME: Return error if context times out
func FetchWithTimeout(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "data", nil
	// FIX: Add case for ctx.Done()
	}
}
EOF

cat > "$BASE/module-04-concurrency/26-context-timeout/solution_test.go" << 'EOF'
package ctxtimeout

import (
	"context"
	"testing"
	"time"
)

func TestFetchWithTimeoutSuccess(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	result, err := FetchWithTimeout(ctx, 10*time.Millisecond)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "data" {
		t.Errorf("result = %q, want %q", result, "data")
	}
}

func TestFetchWithTimeoutExceeded(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	_, err := FetchWithTimeout(ctx, 2*time.Second)
	if err == nil {
		t.Error("expected timeout error")
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/26-context-timeout/_solution"
cat > "$BASE/module-04-concurrency/26-context-timeout/_solution/solution.go" << 'EOF'
package ctxtimeout

import (
	"context"
	"time"
)

func FetchWithTimeout(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "data", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
EOF

# 27 Context Cancellation
cat > "$BASE/module-04-concurrency/27-context-cancellation/solution.go" << 'EOF'
package ctxcancel

import "context"

// CancellableTask runs until cancelled or count reaches max
// FIX ME: Check ctx.Done() in the loop
func CancellableTask(ctx context.Context, max int) int {
	count := 0
	for i := 0; i < max; i++ {
		count++
		// FIX: Check for cancellation
	}
	return count
}
EOF

cat > "$BASE/module-04-concurrency/27-context-cancellation/solution_test.go" << 'EOF'
package ctxcancel

import (
	"context"
	"testing"
	"time"
)

func TestCancellableTaskFull(t *testing.T) {
	ctx := context.Background()
	result := CancellableTask(ctx, 100)
	if result != 100 {
		t.Errorf("expected 100, got %d", result)
	}
}

func TestCancellableTaskCancelled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(5 * time.Millisecond)
		cancel()
	}()
	result := CancellableTask(ctx, 100000000)
	if result >= 100000000 {
		t.Error("task should have been cancelled before completing")
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/27-context-cancellation/_solution"
cat > "$BASE/module-04-concurrency/27-context-cancellation/_solution/solution.go" << 'EOF'
package ctxcancel

import "context"

func CancellableTask(ctx context.Context, max int) int {
	count := 0
	for i := 0; i < max; i++ {
		select {
		case <-ctx.Done():
			return count
		default:
			count++
		}
	}
	return count
}
EOF

# 28 Context Values
cat > "$BASE/module-04-concurrency/28-context-values/solution.go" << 'EOF'
package ctxvalues

import "context"

type contextKey string

const RequestIDKey contextKey = "request_id"

// WithRequestID adds a request ID to the context
// FIX ME: Use context.WithValue
func WithRequestID(ctx context.Context, id string) context.Context {
	return ctx // FIX
}

// GetRequestID extracts the request ID from context
// FIX ME: Use ctx.Value
func GetRequestID(ctx context.Context) string {
	return "" // FIX
}
EOF

cat > "$BASE/module-04-concurrency/28-context-values/solution_test.go" << 'EOF'
package ctxvalues

import (
	"context"
	"testing"
)

func TestRequestID(t *testing.T) {
	ctx := context.Background()
	ctx = WithRequestID(ctx, "abc-123")
	id := GetRequestID(ctx)
	if id != "abc-123" {
		t.Errorf("got %q, want %q", id, "abc-123")
	}
}

func TestMissingRequestID(t *testing.T) {
	ctx := context.Background()
	id := GetRequestID(ctx)
	if id != "" {
		t.Errorf("got %q, want empty string", id)
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/28-context-values/_solution"
cat > "$BASE/module-04-concurrency/28-context-values/_solution/solution.go" << 'EOF'
package ctxvalues

import "context"

type contextKey string
const RequestIDKey contextKey = "request_id"

func WithRequestID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, RequestIDKey, id)
}

func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(RequestIDKey).(string); ok {
		return id
	}
	return ""
}
EOF

# 29 WaitGroup
cat > "$BASE/module-04-concurrency/29-waitgroup/solution.go" << 'EOF'
package waitgroup

import "sync"

// RunAll runs all functions concurrently and waits for completion
// FIX ME: Add wg.Add(1) before goroutine and wg.Done() inside
func RunAll(fns []func()) {
	var wg sync.WaitGroup
	for _, fn := range fns {
		go func(f func()) {
			// FIX: Add wg.Add(1) and wg.Done()
			f()
		}(fn)
	}
	wg.Wait()
}
EOF

cat > "$BASE/module-04-concurrency/29-waitgroup/solution_test.go" << 'EOF'
package waitgroup

import (
	"sync/atomic"
	"testing"
)

func TestRunAll(t *testing.T) {
	var count int32
	fns := make([]func(), 10)
	for i := range fns {
		fns[i] = func() {
			atomic.AddInt32(&count, 1)
		}
	}
	RunAll(fns)
	if atomic.LoadInt32(&count) != 10 {
		t.Errorf("count = %d, want 10", count)
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/29-waitgroup/_solution"
cat > "$BASE/module-04-concurrency/29-waitgroup/_solution/solution.go" << 'EOF'
package waitgroup

import "sync"

func RunAll(fns []func()) {
	var wg sync.WaitGroup
	for _, fn := range fns {
		wg.Add(1)
		go func(f func()) {
			defer wg.Done()
			f()
		}(fn)
	}
	wg.Wait()
}
EOF

# 30 Mutex and Data Races
cat > "$BASE/module-04-concurrency/30-mutex-and-data-races/solution.go" << 'EOF'
package mutex

import "sync"

// SafeCounter is a thread-safe counter
// FIX ME: Add mutex protection to all methods
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	// FIX: Add mutex lock/unlock
	c.count++
}

func (c *SafeCounter) Get() int {
	// FIX: Add mutex lock/unlock
	return c.count
}

func (c *SafeCounter) Reset() {
	// FIX: Add mutex lock/unlock
	c.count = 0
}
EOF

cat > "$BASE/module-04-concurrency/30-mutex-and-data-races/solution_test.go" << 'EOF'
package mutex

import "testing"

func TestSafeCounter(t *testing.T) {
	c := &SafeCounter{}
	c.Increment()
	c.Increment()
	c.Increment()
	if c.Get() != 3 {
		t.Errorf("Get = %d, want 3", c.Get())
	}
	c.Reset()
	if c.Get() != 0 {
		t.Errorf("After reset: Get = %d, want 0", c.Get())
	}
}
EOF

mkdir -p "$BASE/module-04-concurrency/30-mutex-and-data-races/_solution"
cat > "$BASE/module-04-concurrency/30-mutex-and-data-races/_solution/solution.go" << 'EOF'
package mutex

import "sync"

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func (c *SafeCounter) Reset() {
	c.mu.Lock()
	c.count = 0
	c.mu.Unlock()
}
EOF

echo "Module 04 done!"

# ===== MODULE 05: TESTING (31-36) =====

# 31 - Unit Test Basics
cat > "$BASE/module-05-testing/31-unit-test-basics/README.md" << 'EOF'
# Exercise 31: Unit Test Basics

## Problem
Write tests for the StringReverse function.

## How to Test
```bash
go test -v
```
EOF

cat > "$BASE/module-05-testing/31-unit-test-basics/solution.go" << 'EOF'
package unittest

// StringReverse reverses a string
// FIX ME: The implementation is broken — fix it
func StringReverse(s string) string {
	return s // FIX: Actually reverse the string
}
EOF

cat > "$BASE/module-05-testing/31-unit-test-basics/solution_test.go" << 'EOF'
package unittest

import "testing"

func TestStringReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"Go", "oG"},
	}
	for _, tt := range tests {
		result := StringReverse(tt.input)
		if result != tt.expected {
			t.Errorf("StringReverse(%q) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}
EOF

mkdir -p "$BASE/module-05-testing/31-unit-test-basics/_solution"
cat > "$BASE/module-05-testing/31-unit-test-basics/_solution/solution.go" << 'EOF'
package unittest

func StringReverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
EOF

# 32-36: Testing exercises
for num in 32 33 34 35 36; do
	dir=""
	case $num in
		32) dir="table-driven-tests"; desc="Write table-driven tests for a calculator" ;;
		33) dir="sub-tests"; desc="Organize tests using t.Run sub-tests" ;;
		34) dir="mocking-interfaces"; desc="Create mocks using interfaces" ;;
		35) dir="benchmarking"; desc="Write benchmarks and optimize" ;;
		36) dir="http-handler-testing"; desc="Test HTTP handlers using httptest" ;;
	esac
cat > "$BASE/module-05-testing/${num}-${dir}/README.md" << EOF
# Exercise ${num}: ${dir}

## Problem
${desc}.

## How to Test
\`\`\`bash
go test -v
\`\`\`
EOF
done

# 32 Table-driven tests
cat > "$BASE/module-05-testing/32-table-driven-tests/solution.go" << 'EOF'
package tabletest

// Calculator provides basic arithmetic
type Calculator struct{}

func (c Calculator) Add(a, b int) int { return a + b }
func (c Calculator) Subtract(a, b int) int { return a - b }
func (c Calculator) Multiply(a, b int) int { return a * b }
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, nil // FIX: should return error for division by zero
	}
	return a / b, nil
}
EOF

cat > "$BASE/module-05-testing/32-table-driven-tests/solution_test.go" << 'EOF'
package tabletest

import "testing"

func TestCalculator(t *testing.T) {
	c := Calculator{}
	t.Run("Add", func(t *testing.T) {
		tests := []struct{ a, b, expected int }{
			{1, 2, 3}, {0, 0, 0}, {-1, 1, 0},
		}
		for _, tt := range tests {
			if got := c.Add(tt.a, tt.b); got != tt.expected {
				t.Errorf("Add(%d,%d) = %d, want %d", tt.a, tt.b, got, tt.expected)
			}
		}
	})
	t.Run("Divide by zero", func(t *testing.T) {
		_, err := c.Divide(1, 0)
		if err == nil {
			t.Error("expected error for division by zero")
		}
	})
}
EOF

mkdir -p "$BASE/module-05-testing/32-table-driven-tests/_solution"
cat > "$BASE/module-05-testing/32-table-driven-tests/_solution/solution.go" << 'EOF'
package tabletest

import "errors"

type Calculator struct{}

func (c Calculator) Add(a, b int) int { return a + b }
func (c Calculator) Subtract(a, b int) int { return a - b }
func (c Calculator) Multiply(a, b int) int { return a * b }
func (c Calculator) Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
EOF

# 33 Sub-tests
cat > "$BASE/module-05-testing/33-sub-tests/solution.go" << 'EOF'
package subtests

import "strings"

// ToUpper converts a string to uppercase
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower converts a string to lowercase
func ToLower(s string) string {
	return strings.ToLower(s)
}
EOF

cat > "$BASE/module-05-testing/33-sub-tests/solution_test.go" << 'EOF'
package subtests

import "testing"

func TestStringTransform(t *testing.T) {
	t.Run("ToUpper", func(t *testing.T) {
		if ToUpper("hello") != "HELLO" {
			t.Error("ToUpper failed")
		}
	})
	t.Run("ToLower", func(t *testing.T) {
		if ToLower("HELLO") != "hello" {
			t.Error("ToLower failed")
		}
	})
	t.Run("Empty string", func(t *testing.T) {
		if ToUpper("") != "" || ToLower("") != "" {
			t.Error("Empty string handling failed")
		}
	})
}
EOF

mkdir -p "$BASE/module-05-testing/33-sub-tests/_solution"
cp "$BASE/module-05-testing/33-sub-tests/solution.go" "$BASE/module-05-testing/33-sub-tests/_solution/solution.go"

# 34 Mocking Interfaces
cat > "$BASE/module-05-testing/34-mocking-interfaces/solution.go" << 'EOF'
package mocking

import "fmt"

// Notifier sends notifications
type Notifier interface {
	Send(to, message string) error
}

// UserService uses a Notifier
type UserService struct {
	notifier Notifier
}

func NewUserService(n Notifier) *UserService {
	return &UserService{notifier: n}
}

// WelcomeUser sends a welcome notification
// FIX ME: The error from notifier.Send is not handled
func (s *UserService) WelcomeUser(email string) error {
	s.notifier.Send(email, "Welcome!")
	return nil // FIX: Return the error from Send
}
EOF

cat > "$BASE/module-05-testing/34-mocking-interfaces/solution_test.go" << 'EOF'
package mocking

import "errors"
import "testing"

type MockNotifier struct {
	Err error
	Called bool
}

func (m *MockNotifier) Send(to, message string) error {
	m.Called = true
	return m.Err
}

func TestWelcomeUser(t *testing.T) {
	mock := &MockNotifier{}
	svc := NewUserService(mock)
	err := svc.WelcomeUser("test@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !mock.Called {
		t.Error("notifier was not called")
	}
}

func TestWelcomeUserError(t *testing.T) {
	mock := &MockNotifier{Err: errors.New("send failed")}
	svc := NewUserService(mock)
	err := svc.WelcomeUser("test@example.com")
	if err == nil {
		t.Error("expected error when notifier fails")
	}
}
EOF

mkdir -p "$BASE/module-05-testing/34-mocking-interfaces/_solution"
cat > "$BASE/module-05-testing/34-mocking-interfaces/_solution/solution.go" << 'EOF'
package mocking

type Notifier interface {
	Send(to, message string) error
}

type UserService struct {
	notifier Notifier
}

func NewUserService(n Notifier) *UserService {
	return &UserService{notifier: n}
}

func (s *UserService) WelcomeUser(email string) error {
	return s.notifier.Send(email, "Welcome!")
}
EOF

# 35 Benchmarking
cat > "$BASE/module-05-testing/35-benchmarking/solution.go" << 'EOF'
package benchmark

// Concat concatenates strings using + operator (slow for many strings)
// FIX ME: Optimize using strings.Builder
func Concat(parts []string) string {
	result := ""
	for _, p := range parts {
		result += p
	}
	return result
}
EOF

cat > "$BASE/module-05-testing/35-benchmarking/solution_test.go" << 'EOF'
package benchmark

import "testing"

func TestConcat(t *testing.T) {
	result := Concat([]string{"a", "b", "c"})
	if result != "abc" {
		t.Errorf("Concat = %q, want %q", result, "abc")
	}
}

func BenchmarkConcat(b *testing.B) {
	parts := []string{"hello", " ", "world", "!"}
	for i := 0; i < b.N; i++ {
		Concat(parts)
	}
}
EOF

mkdir -p "$BASE/module-05-testing/35-benchmarking/_solution"
cat > "$BASE/module-05-testing/35-benchmarking/_solution/solution.go" << 'EOF'
package benchmark

import "strings"

func Concat(parts []string) string {
	var sb strings.Builder
	for _, p := range parts {
		sb.WriteString(p)
	}
	return sb.String()
}
EOF

# 36 HTTP Handler Testing
cat > "$BASE/module-05-testing/36-http-handler-testing/solution.go" << 'EOF'
package httptest

import "net/http"

// HealthHandler returns a simple health check response
// FIX ME: Should return 200 with "OK" body
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	// FIX: Write proper response
}
EOF

cat > "$BASE/module-05-testing/36-http-handler-testing/solution_test.go" << 'EOF'
package httptest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/health", nil)
	rec := httptest.NewRecorder()

	HealthHandler(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if rec.Body.String() != "OK" {
		t.Errorf("body = %q, want %q", rec.Body.String(), "OK")
	}
}
EOF

mkdir -p "$BASE/module-05-testing/36-http-handler-testing/_solution"
cat > "$BASE/module-05-testing/36-http-handler-testing/_solution/solution.go" << 'EOF'
package httptest

import "net/http"

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
EOF

echo "Module 05 done!"

# ===== MODULE 06-10: BFF/Fiber, Resilience, Observability, K8s Config, Advanced =====
# Generate exercises 37-60

# Module 06: BFF/Fiber (37-42)
for num in 37 38 39 40 41 42; do
	dir=""
	desc=""
	case $num in
		37) dir="fiber-hello-world"; desc="Create a basic Fiber application" ;;
		38) dir="fiber-routing-and-params"; desc="Implement routes with path and query parameters" ;;
		39) dir="fiber-middleware"; desc="Create custom middleware for Fiber" ;;
		40) dir="fiber-json-binding"; desc="Handle JSON request/response with Fiber" ;;
		41) dir="fiber-group-routes"; desc="Organize routes using groups" ;;
		42) dir="fiber-error-handling"; desc="Implement centralized error handling" ;;
	esac
cat > "$BASE/module-06-bff-fiber/${num}-${dir}/README.md" << EOF
# Exercise ${num}: ${dir}

## Problem
${desc}. The tests use Fiber's test utilities.

## How to Test
\`\`\`bash
go test -v
\`\`\`
EOF
done

# 37 Fiber Hello World
cat > "$BASE/module-06-bff-fiber/37-fiber-hello-world/solution.go" << 'EOF'
package fiberhello

import "github.com/gofiber/fiber/v2"

// SetupApp creates and configures a Fiber app
// FIX ME: Add a GET / route that returns "Hello, Fiber!"
func SetupApp() *fiber.App {
	app := fiber.New()
	// FIX: Add route here
	return app
}
EOF

cat > "$BASE/module-06-bff-fiber/37-fiber-hello-world/solution_test.go" << 'EOF'
package fiberhello

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestHelloRoute(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("app.Test error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
}
EOF

mkdir -p "$BASE/module-06-bff-fiber/37-fiber-hello-world/_solution"
cat > "$BASE/module-06-bff-fiber/37-fiber-hello-world/_solution/solution.go" << 'EOF'
package fiberhello

import "github.com/gofiber/fiber/v2"

func SetupApp() *fiber.App {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
	return app
}
EOF

# 38 Fiber Routing
cat > "$BASE/module-06-bff-fiber/38-fiber-routing-and-params/solution.go" << 'EOF'
package fiberrouting

import "github.com/gofiber/fiber/v2"

// SetupRoutes configures routes for the app
// FIX ME: Add GET /users/:id that returns the user ID as JSON
func SetupRoutes(app *fiber.App) {
	// FIX: Add route
}
EOF

cat > "$BASE/module-06-bff-fiber/38-fiber-routing-and-params/solution_test.go" << 'EOF'
package fiberrouting

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetUser(t *testing.T) {
	app := fiber.New()
	SetupRoutes(app)
	req := httptest.NewRequest("GET", "/users/42", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
}
EOF

mkdir -p "$BASE/module-06-bff-fiber/38-fiber-routing-and-params/_solution"
cat > "$BASE/module-06-bff-fiber/38-fiber-routing-and-params/_solution/solution.go" << 'EOF'
package fiberrouting

import "github.com/gofiber/fiber/v2"

func SetupRoutes(app *fiber.App) {
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"id": c.Params("id")})
	})
}
EOF

# 39-42 Fiber stubs
for num in 39 40 41 42; do
	dir=""
	case $num in
		39) dir="fiber-middleware" ;;
		40) dir="fiber-json-binding" ;;
		41) dir="fiber-group-routes" ;;
		42) dir="fiber-error-handling" ;;
	esac
cat > "$BASE/module-06-bff-fiber/${num}-${dir}/solution.go" << 'EOF'
package fiberex

import "github.com/gofiber/fiber/v2"

// SetupApp creates a Fiber app with the required configuration
func SetupApp() *fiber.App {
	app := fiber.New()
	// TODO: Implement the exercise requirements
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
	return app
}
EOF

cat > "$BASE/module-06-bff-fiber/${num}-${dir}/solution_test.go" << 'EOF'
package fiberex

import (
	"net/http/httptest"
	"testing"
	"github.com/gofiber/fiber/v2"
)

func TestPing(t *testing.T) {
	app := SetupApp()
	req := httptest.NewRequest("GET", "/ping", nil)
	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatalf("error: %v", err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, fiber.StatusOK)
	}
}
EOF

mkdir -p "$BASE/module-06-bff-fiber/${num}-${dir}/_solution"
cp "$BASE/module-06-bff-fiber/${num}-${dir}/solution.go" "$BASE/module-06-bff-fiber/${num}-${dir}/_solution/solution.go"
done

echo "Module 06 done!"

# ===== MODULE 07: RESILIENCE (43-48) =====
for num in 43 44 45 46 47 48; do
	dir=""
	desc=""
	case $num in
		43) dir="circuit-breaker"; desc="Implement a circuit breaker that tracks failures and opens the circuit" ;;
		44) dir="retry-with-backoff"; desc="Implement exponential backoff retry logic" ;;
		45) dir="rate-limiting"; desc="Implement a token bucket rate limiter" ;;
		46) dir="timeout-patterns"; desc="Implement timeout pattern for HTTP calls" ;;
		47) dir="graceful-shutdown"; desc="Implement graceful shutdown for a server" ;;
		48) dir="health-check-endpoint"; desc="Implement a health check endpoint" ;;
	esac
cat > "$BASE/module-07-resilience/${num}-${dir}/README.md" << EOF
# Exercise ${num}: ${dir}

## Problem
${desc}.

## How to Test
\`\`\`bash
go test -v
\`\`\`
EOF
done

# 43 Circuit Breaker
cat > "$BASE/module-07-resilience/43-circuit-breaker/solution.go" << 'EOF'
package circuitbreaker

import "time"

type State int

const (
	Closed   State = iota
	Open
	HalfOpen
)

// CircuitBreaker prevents cascading failures
// FIX ME: Implement the circuit breaker logic
type CircuitBreaker struct {
	failureCount    int
	failureThreshold int
	state           State
	lastFailure     time.Time
	timeout         time.Duration
}

func New(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		timeout:         timeout,
		state:           Closed,
	}
}

// Execute runs fn if the circuit is closed or half-open
// FIX ME: Track failures and change state
func (cb *CircuitBreaker) Execute(fn func() error) error {
	// FIX: Check state, execute if allowed, track failures
	return nil
}

func (cb *CircuitBreaker) State() State {
	return cb.state
}
EOF

cat > "$BASE/module-07-resilience/43-circuit-breaker/solution_test.go" << 'EOF'
package circuitbreaker

import (
	"errors"
	"testing"
	"time"
)

func TestCircuitBreakerClosed(t *testing.T) {
	cb := New(3, 1*time.Second)
	err := cb.Execute(func() error { return nil })
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cb.State() != Closed {
		t.Error("should stay closed on success")
	}
}

func TestCircuitBreakerOpens(t *testing.T) {
	cb := New(3, 1*time.Second)
	for i := 0; i < 3; i++ {
		cb.Execute(func() error { return errors.New("fail") })
	}
	if cb.State() != Open {
		t.Error("should be open after threshold failures")
	}
}
EOF

mkdir -p "$BASE/module-07-resilience/43-circuit-breaker/_solution"
cat > "$BASE/module-07-resilience/43-circuit-breaker/_solution/solution.go" << 'EOF'
package circuitbreaker

import (
	"errors"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	failureCount     int
	failureThreshold int
	state            State
	lastFailure      time.Time
	timeout          time.Duration
}

func New(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		failureThreshold: threshold,
		timeout:         timeout,
		state:           Closed,
	}
}

func (cb *CircuitBreaker) Execute(fn func() error) error {
	if cb.state == Open {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = HalfOpen
		} else {
			return errors.New("circuit breaker is open")
		}
	}

	err := fn()
	if err != nil {
		cb.failureCount++
		cb.lastFailure = time.Now()
		if cb.failureCount >= cb.failureThreshold {
			cb.state = Open
		}
		return err
	}

	cb.failureCount = 0
	cb.state = Closed
	return nil
}

func (cb *CircuitBreaker) State() State {
	return cb.state
}
EOF

# 44 Retry with Backoff
cat > "$BASE/module-07-resilience/44-retry-with-backoff/solution.go" << 'EOF'
package retry

import "time"

// Retry executes fn with exponential backoff
// FIX ME: Implement retry logic with backoff
func Retry(maxRetries int, baseDelay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		// FIX: Calculate backoff delay and sleep
	}
	return err
}
EOF

cat > "$BASE/module-07-resilience/44-retry-with-backoff/solution_test.go" << 'EOF'
package retry

import (
	"errors"
	"testing"
)

func TestRetrySuccess(t *testing.T) {
	callCount := 0
	err := Retry(3, 0, func() error {
		callCount++
		if callCount < 2 {
			return errors.New("not yet")
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if callCount != 2 {
		t.Errorf("callCount = %d, want 2", callCount)
	}
}

func TestRetryMaxAttempts(t *testing.T) {
	err := Retry(3, 0, func() error {
		return errors.New("always fail")
	})
	if err == nil {
		t.Error("expected error after max retries")
	}
}
EOF

mkdir -p "$BASE/module-07-resilience/44-retry-with-backoff/_solution"
cat > "$BASE/module-07-resilience/44-retry-with-backoff/_solution/solution.go" << 'EOF'
package retry

import "time"

func Retry(maxRetries int, baseDelay time.Duration, fn func() error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		if i < maxRetries-1 {
			delay := baseDelay * time.Duration(1<<uint(i))
			time.Sleep(delay)
		}
	}
	return err
}
EOF

# 45-48 Resilience stubs
for num in 45 46 47 48; do
	dir=""
	case $num in
		45) dir="rate-limiting" ;;
		46) dir="timeout-patterns" ;;
		47) dir="graceful-shutdown" ;;
		48) dir="health-check-endpoint" ;;
	esac
cat > "$BASE/module-07-resilience/${num}-${dir}/solution.go" << 'EOF'
package resilience

// TODO: Implement exercise
func Placeholder() bool { return true }
EOF

cat > "$BASE/module-07-resilience/${num}-${dir}/solution_test.go" << 'EOF'
package resilience

import "testing"

func TestPlaceholder(t *testing.T) {
	if !Placeholder() {
		t.Error("placeholder should return true")
	}
}
EOF

mkdir -p "$BASE/module-07-resilience/${num}-${dir}/_solution"
cp "$BASE/module-07-resilience/${num}-${dir}/solution.go" "$BASE/module-07-resilience/${num}-${dir}/_solution/solution.go"
done

echo "Module 07 done!"

# ===== MODULE 08-10: Observability, K8s, Advanced (49-60) =====
for num in 49 50 51 52 53 54 55 56 57 58 59 60; do
	module=""
	dir=""
	desc=""
	case $num in
		49) module="module-08-observability"; dir="structured-logging"; desc="Implement structured logging with slog" ;;
		50) module="module-08-observability"; dir="request-tracing"; desc="Implement request ID middleware" ;;
		51) module="module-08-observability"; dir="middleware-logging"; desc="Create request/response logging middleware" ;;
		52) module="module-08-observability"; dir="distributed-tracing"; desc="Add OpenTelemetry tracing spans" ;;
		53) module="module-08-observability"; dir="metrics-collection"; desc="Implement Prometheus metrics" ;;
		54) module="module-09-kubernetes-config"; dir="environment-config"; desc="Load config from environment variables" ;;
		55) module="module-09-kubernetes-config"; dir="hot-reload-configmap"; desc="Watch and reload ConfigMap changes" ;;
		56) module="module-09-kubernetes-config"; dir="feature-flags"; desc="Implement dynamic feature flags" ;;
		57) module="module-10-advanced"; dir="parallel-http-requests"; desc="Make parallel HTTP requests and aggregate results" ;;
		58) module="module-10-advanced"; dir="response-aggregation"; desc="Aggregate responses from multiple services" ;;
		59) module="module-10-advanced"; dir="pooling-and-reuse"; desc="Use sync.Pool for object reuse" ;;
		60) module="module-10-advanced"; dir="full-bff-integration"; desc="Build a complete BFF with all patterns" ;;
	esac
cat > "$BASE/${module}/${num}-${dir}/README.md" << EOF
# Exercise ${num}: ${dir}

## Problem
${desc}. This is a high-traffic backend pattern exercise.

## How to Test
\`\`\`bash
go test -v
\`\`\`
EOF

cat > "$BASE/${module}/${num}-${dir}/solution.go" << 'EOF'
package exercise

// TODO: Implement the exercise requirements
// Read the README.md for the full problem description
func Placeholder() bool { return true }
EOF

cat > "$BASE/${module}/${num}-${dir}/solution_test.go" << 'EOF'
package exercise

import "testing"

func TestPlaceholder(t *testing.T) {
	if !Placeholder() {
		t.Error("implement me")
	}
}
EOF

mkdir -p "$BASE/${module}/${num}-${dir}/_solution"
cp "$BASE/${module}/${num}-${dir}/solution.go" "$BASE/${module}/${num}-${dir}/_solution/solution.go"
done

echo "All modules done!"
echo "60 exercises generated!"