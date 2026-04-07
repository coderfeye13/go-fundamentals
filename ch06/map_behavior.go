package main

import "fmt"

// MAPS:
// Maps are passed by value, but the value contains a pointer internally.
// So modifying the map affects the original data.

func updateMap(m map[string]int) {
	// Modifying map affects original
	m["a"] = 100
}

func replaceMap(m map[string]int) {
	// Reassigning map only changes the local copy
	m = map[string]int{
		"new": 999,
	}

	fmt.Println("Inside replaceMap:", m)
}

func main() {
	fmt.Println("=== Map Behavior ===")

	m := map[string]int{
		"a": 1,
		"b": 2,
	}

	fmt.Println("Original:", m)

	updateMap(m)
	fmt.Println("After updateMap:", m)

	replaceMap(m)
	fmt.Println("After replaceMap:", m)
}
