package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// new() allocates memory, sets zero value, returns a pointer.
	// Fields must be set after creation.
	p1 := new(int)
	fmt.Printf("new(int)    → value: %d | type: %T\n", *p1, p1)

	p2 := new(Point)
	fmt.Printf("new(Point)  → value: %v | type: %T\n", *p2, p2)
	p2.X = 10
	p2.Y = 20

	// & with struct literal is more idiomatic.
	// Allows inline initialization.
	p3 := &Point{X: 10, Y: 20}
	fmt.Printf("&Point{...} → value: %v | type: %T\n", *p3, p3)

	// Both return the same type (*Point).
	// Allocation (stack vs heap) is decided by the compiler.
	fmt.Println("\nBoth p2 and p3 are of type *Point")

	// new() is useful for primitives when you need a pointer to zero value.
	flag := new(bool)
	fmt.Printf("\nnew(bool) → value: %v | is nil: %v\n", *flag, flag == nil)

	// A declared pointer without initialization is nil.
	var nilBool *bool
	fmt.Printf("var *bool  → is nil: %v\n", nilBool == nil)
}
