package main

import "fmt"

// --- BY VALUE ---
// Go passes a COPY of the int into this function.
// The original variable is never touched.
func doubleByValue(n int) {
	n *= 2
	fmt.Printf("  [inside doubleByValue] n = %d\n", n)
}

// --- BY POINTER ---
// Go passes the memory address.
// We dereference (*n) to reach and modify the original value.
func doubleByPointer(n *int) {
	*n *= 2
	fmt.Printf("  [inside doubleByPointer] *n = %d\n", *n)
}

// --- STRUCT: BY VALUE ---
// Entire struct is copied. Expensive for large structs.
// Changes do NOT affect the original.
type Config struct {
	Timeout int
	Retries int
}

func resetTimeoutByValue(c Config) {
	c.Timeout = 0 // modifies the copy only
}

// --- STRUCT: BY POINTER ---
// Only the address (8 bytes) is copied. Efficient.
// Changes DO affect the original.
func resetTimeoutByPointer(c *Config) {
	c.Timeout = 0 // modifies the original
}

func main() {
	fmt.Println("=== Primitive: Value vs Pointer ===")

	x := 10
	fmt.Printf("Before doubleByValue:   x = %d\n", x)
	doubleByValue(x)
	fmt.Printf("After  doubleByValue:   x = %d\n\n", x) // still 10

	fmt.Printf("Before doubleByPointer: x = %d\n", x)
	doubleByPointer(&x)
	fmt.Printf("After  doubleByPointer: x = %d\n\n", x) // now 20

	fmt.Println("=== Struct: Value vs Pointer ===")

	cfg := Config{Timeout: 30, Retries: 3}

	fmt.Printf("Before resetByValue:   Timeout = %d\n", cfg.Timeout)
	resetTimeoutByValue(cfg)
	fmt.Printf("After  resetByValue:   Timeout = %d\n\n", cfg.Timeout) // still 30

	fmt.Printf("Before resetByPointer: Timeout = %d\n", cfg.Timeout)
	resetTimeoutByPointer(&cfg)
	fmt.Printf("After  resetByPointer: Timeout = %d\n\n", cfg.Timeout) // now 0
}
