package main

import "fmt"

func main() {
	// 1. Declaration
	// ptr is a pointer to an int, but it has no address yet.
	// Its zero value is 'nil'.
	var ptr *int

	fmt.Println("Is ptr nil?", ptr == nil) // Prints: true

	// DANGER ZONE (TEHLİKELİ BÖLGE)
	// The following line would cause a CRASH (panic) because
	// we are trying to go to a 'nil' address.

	// fmt.Println(*ptr) // <--- UNCOMMENT THIS TO SEE THE PANIC

	// SAFE ZONE (GÜVENLİ BÖLGE - THE SOLUTION)
	// This is called a "Nil Check". It's a standard pattern
	// in Go (and DevOps tools) to prevent crashes.

	if ptr != nil {
		// This block only runs if ptr points to a real memory address.
		fmt.Printf("The value is: %d\n", *ptr)
	} else {
		// This block handles the case where there is no data.
		fmt.Println("Pointer is nil, skipping dereference to avoid a crash!")
	}

	// EXAMPLE WITH DATA (VERİ İLE ÖRNEK)
	x := 100
	ptr = &x // Now ptr points to x's address

	if ptr != nil {
		fmt.Printf("Now ptr is NOT nil. Value: %d\n", *ptr) // Prints: 100
	}
}
