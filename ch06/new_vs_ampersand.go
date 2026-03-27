package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	// --- new() ---
	// Allocates memory for the type, sets it to zero value, returns a pointer.
	// You cannot initialize fields inline with new().
	p1 := new(int)
	fmt.Printf("new(int)    → value: %d | type: %T\n", *p1, p1) // 0, *int

	p2 := new(Point)
	fmt.Printf("new(Point)  → value: %v | type: %T\n", *p2, p2) // {0 0}, *Point
	p2.X = 10                                                   // set fields after the fact
	p2.Y = 20

	// --- & (address-of) with struct literal ---
	// More idiomatic Go. Lets you initialize fields inline.
	p3 := &Point{X: 10, Y: 20}
	fmt.Printf("&Point{...} → value: %v | type: %T\n", *p3, p3) // {10 20}, *Point

	// Both result in the same type (*Point), both live on the heap.
	// In practice: prefer &T{} over new(T) for structs.
	fmt.Println("\nAre p2 and p3 the same type?", fmt.Sprintf("%T", p2) == fmt.Sprintf("%T", p3)) // true

	// --- new() is sometimes useful for primitives ---
	// e.g. when you need a *bool or *int with a zero value pointer
	flag := new(bool)
	fmt.Printf("\nnew(bool) → value: %v | is nil: %v\n", *flag, flag == nil) // false, false

	// Compare: var b *bool → this is nil, points to nothing
	var nilBool *bool
	fmt.Printf("var *bool  → is nil: %v\n", nilBool == nil) // true
}
