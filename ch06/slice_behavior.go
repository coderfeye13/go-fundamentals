package main

import "fmt"

// SLICES:
// A slice contains:
// - pointer to underlying array
// - length
// - capacity

func modifyElement(s []int) {
	// Modifying an element affects the original slice
	s[0] = 100
}

func appendElement(s []int) {
	// Append updates only the local slice header.
	// The original slice in main keeps its old length.
	s = append(s, 999)

	fmt.Println("Inside appendElement:", s, "len:", len(s), "cap:", cap(s))
}

func main() {
	fmt.Println("=== Slice Behavior ===")

	// len = 2, cap = 4
	s := make([]int, 2, 4)
	s[0] = 1
	s[1] = 2

	fmt.Println("Original:", s, "len:", len(s), "cap:", cap(s))

	modifyElement(s)
	fmt.Println("After modifyElement:", s, "len:", len(s), "cap:", cap(s))

	appendElement(s)
	fmt.Println("After appendElement:", s, "len:", len(s), "cap:", cap(s))

	// A sub-slice shares the same underlying array
	sub := s[:1]
	sub[0] = 500

	fmt.Println("After modifying sub-slice:", s)
}
