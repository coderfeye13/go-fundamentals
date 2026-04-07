package main

import "fmt"

func main() {
	// ptr is a pointer to int.
	// Its zero value is nil.
	var ptr *int

	fmt.Println("Is ptr nil?", ptr == nil)

	// Uncomment to trigger a panic:
	// fmt.Println(*ptr)

	// Safe dereference with nil check.
	if ptr != nil {
		fmt.Printf("Value: %d\n", *ptr)
	} else {
		fmt.Println("Pointer is nil, skipping dereference.")
	}

	// ptr now points to a real value.
	x := 100
	ptr = &x

	if ptr != nil {
		fmt.Printf("Now ptr is not nil. Value: %d\n", *ptr)
	}
}
