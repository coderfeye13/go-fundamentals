package main

import (
	"errors"
	"fmt"
)

func main() {
	// Exercise 3: variable shadowing in a loop
	total := 0
	for i := 1; i < 10; i++ {
		total := total + i
		fmt.Println("inner total:", total)
	}
	fmt.Println("outer total:", total)

	shadowInIf()
	shadowWithMultipleAssignment()
	errShadowingBug()
}

// 1) Shadowing inside an if block
func shadowInIf() {
	fmt.Println("\nShadowing inside if block:")

	x := 10

	if x > 5 {
		fmt.Println("before shadowing, outer x:", x)
		x := 20
		fmt.Println("inner x:", x)
	}

	fmt.Println("outer x after if:", x)
}

// 2) Shadowing with := and multiple assignment
func shadowWithMultipleAssignment() {
	fmt.Println("\nShadowing with multiple assignment:")

	x := 10

	if x > 5 {
		x, y := 5, 20
		fmt.Println("inner x and y:", x, y)
	}

	fmt.Println("outer x after block:", x)
}

// errShadowingBug shows how using := inside a block can create a new err
// variable that shadows the outer err, leading to subtle bugs.
func errShadowingBug() {
	fmt.Println("Error shadowing example:")

	var err error

	if true {
		err := errors.New("inner error happened")
		fmt.Println("inside if:", err)
	}

	if err != nil {
		fmt.Println("outer err is not nil:", err)
	} else {
		fmt.Println("outer err is still nil")
	}
}
