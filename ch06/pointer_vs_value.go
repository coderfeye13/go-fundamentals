package main

import "fmt"

// Go is always pass-by-value.
// For primitives, passing by value means the original variable is unchanged.

func doubleByValue(n int) {
	n *= 2
	fmt.Printf("  [inside doubleByValue] n = %d\n", n)
}

// When passing a pointer, Go still copies a value,
// but that value is the pointer (address).
// Dereferencing lets us modify the original value.

func doubleByPointer(n *int) {
	*n *= 2
	fmt.Printf("  [inside doubleByPointer] *n = %d\n", *n)
}

// Struct passed by value:
// the function receives a copy,
// so changes do not affect the original.

type Config struct {
	Timeout int
	Retries int
}

func resetTimeoutByValue(c Config) {
	c.Timeout = 0
}

// Struct passed by pointer:
// the function receives a copy of the pointer,
// and can modify the original struct through it.

func resetTimeoutByPointer(c *Config) {
	c.Timeout = 0
}

func main() {
	fmt.Println("=== Primitive: Value vs Pointer ===")

	x := 10
	fmt.Printf("Before doubleByValue:   x = %d\n", x)
	doubleByValue(x)
	fmt.Printf("After  doubleByValue:   x = %d\n\n", x)

	fmt.Printf("Before doubleByPointer: x = %d\n", x)
	doubleByPointer(&x)
	fmt.Printf("After  doubleByPointer: x = %d\n\n", x)

	fmt.Println("=== Struct: Value vs Pointer ===")

	cfg := Config{Timeout: 30, Retries: 3}

	fmt.Printf("Before resetByValue:   Timeout = %d\n", cfg.Timeout)
	resetTimeoutByValue(cfg)
	fmt.Printf("After  resetByValue:   Timeout = %d\n\n", cfg.Timeout)

	fmt.Printf("Before resetByPointer: Timeout = %d\n", cfg.Timeout)
	resetTimeoutByPointer(&cfg)
	fmt.Printf("After  resetByPointer: Timeout = %d\n\n", cfg.Timeout)
}
