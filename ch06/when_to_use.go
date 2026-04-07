package main

import "fmt"

// RULE 1: USE POINTER when the function must modify the original value.

type Counter struct {
	Value int
}

func increment(c *Counter) {
	c.Value++
}

// RULE 2: USE VALUE for small, simple, read-only data.

func isEven(n int) bool {
	return n%2 == 0
}

// RULE 3: CONSIDER POINTER for large structs to avoid expensive copying.

type HeavyConfig struct {
	Data    [1024]byte
	Name    string
	Version int
}

// The pointer value is copied instead of the whole struct.
func describeConfig(cfg *HeavyConfig) {
	fmt.Printf("Config: %s v%d\n", cfg.Name, cfg.Version)
}

// RULE 4: USE POINTER when nil has a meaningful meaning.
// Useful for optional fields, especially in JSON/API models.

type User struct {
	Name  string
	Email *string
}

func printEmail(u User) {
	if u.Email == nil {
		fmt.Printf("%s has no email on file.\n", u.Name)
	} else {
		fmt.Printf("%s → %s\n", u.Name, *u.Email)
	}
}

// RULE 5: USE POINTER RECEIVERS when methods modify struct state.

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s Stack) Peek() int {
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
	fmt.Println("Counter:", c.Value)

	fmt.Println("\n=== Rule 2: Value for small read-only data ===")
	fmt.Println("isEven(4):", isEven(4))
	fmt.Println("isEven(7):", isEven(7))

	fmt.Println("\n=== Rule 3: Pointer for large structs ===")
	cfg := &HeavyConfig{Name: "AppConfig", Version: 3}
	describeConfig(cfg)

	fmt.Println("\n=== Rule 4: Pointer for optional/nil state ===")
	email := "alice@example.com"
	u1 := User{Name: "Alice", Email: &email}
	u2 := User{Name: "Bob", Email: nil}
	printEmail(u1)
	printEmail(u2)

	fmt.Println("\n=== Rule 5: Pointer receiver for methods ===")
	s := Stack{}
	s.Push(10)
	s.Push(20)
	s.Push(30)
	fmt.Println("Top of stack:", s.Peek())
}
