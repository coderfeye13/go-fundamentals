package main

import "fmt"

func main() {

	fmt.Println("===== BASIC PROGRAM STRUCTURE =====")

	// Every Go program starts with package main
	// Execution starts from func main()

	fmt.Println("Hello, Go!")

	fmt.Println("\n===== VARIABLES =====")

	// var declaration
	var x int = 10
	fmt.Println("x:", x)

	// type inference
	var y = 20
	fmt.Println("y:", y)

	// short declaration (inside function only)
	z := 30
	fmt.Println("z:", z)

	fmt.Println("\n===== MULTIPLE VARIABLES =====")

	a, b := 1, "hello"
	fmt.Println("a:", a, "b:", b)

	fmt.Println("\n===== ZERO VALUES =====")

	var intZero int
	var boolZero bool
	var stringZero string

	fmt.Println("int zero:", intZero)
	fmt.Println("bool zero:", boolZero)
	fmt.Println("string zero:", stringZero)

	fmt.Println("\n===== UNUSED VARIABLE RULE =====")

	// Uncommenting next line causes compile error:
	// unused := 100

	// Go does not allow unused local variables
}
