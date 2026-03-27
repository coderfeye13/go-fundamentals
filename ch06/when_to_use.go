package main

import "fmt"

// ============================================================
// RULE 1: USE POINTER — when you need to mutate the argument
// ============================================================

type Counter struct {
	Value int
}

func increment(c *Counter) {
	c.Value++
}

// ============================================================
// RULE 2: USE VALUE — for small, simple, read-only data
// Copying an int or small struct is cheaper than indirection.
// ============================================================

func isEven(n int) bool {
	return n%2 == 0
}

// ============================================================
// RULE 3: USE POINTER — for large structs (avoid expensive copy)
// ============================================================

type HeavyConfig struct {
	Data    [1024]byte // 1 KB of data
	Name    string
	Version int
}

// Passing by pointer: only 8 bytes (address) copied, not 1 KB+
func describeConfig(cfg *HeavyConfig) {
	fmt.Printf("Config: %s v%d\n", cfg.Name, cfg.Version)
}

// ============================================================
// RULE 4: USE POINTER — when nil is a valid/meaningful state
// (optional data, missing values, uninitialized state)
// ============================================================

type User struct {
	Name  string
	Email *string // nil means "no email provided" — intentional!
}

func printEmail(u User) {
	if u.Email == nil {
		fmt.Printf("%s has no email on file.\n", u.Name)
	} else {
		fmt.Printf("%s → %s\n", u.Name, *u.Email)
	}
}

// ============================================================
// RULE 5: USE POINTER — for method receivers that modify state
// ============================================================

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) { // pointer receiver → modifies s
	s.items = append(s.items, item)
}

func (s Stack) Peek() int { // value receiver → read-only, no mutation
	if len(s.items) == 0 {
		return -1
	}
	return s.items[len(s.items)-1]
}

func main() {
	fmt.Println("=== Rule 1: Mutate via pointer ===")
	c := Counter{}
	increment(&c)
	increment(&c)
	fmt.Println("Counter:", c.Value) // 2

	fmt.Println("\n=== Rule 2: Value for small read-only data ===")
	fmt.Println("isEven(4):", isEven(4)) // true
	fmt.Println("isEven(7):", isEven(7)) // false

	fmt.Println("\n=== Rule 3: Pointer for large structs ===")
	cfg := &HeavyConfig{Name: "AppConfig", Version: 3}
	describeConfig(cfg)

	fmt.Println("\n=== Rule 4: Pointer for optional/nil state ===")
	email := "alice@example.com"
	u1 := User{Name: "Alice", Email: &email}
	u2 := User{Name: "Bob", Email: nil}
	printEmail(u1) // Alice → alice@example.com
	printEmail(u2) // Bob has no email on file.

	fmt.Println("\n=== Rule 5: Pointer receiver for methods ===")
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Top of stack:", s.Peek()) // 30
}
