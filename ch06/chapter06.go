package main

import "fmt"

func main() {
	var x int = 42
	var ptr *int = &x // ptr is a pointer to an int

	fmt.Printf("x value: %d\n", x)
	fmt.Printf("x address: %v\n", ptr)
	fmt.Printf("Value at address: %d\n", *ptr)

	// Modification via pointer
	*ptr = 100
	fmt.Printf("New x value: %d\n", x) // x is now 100!
}
