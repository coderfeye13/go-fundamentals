package main

import "fmt"

func UpdateSlice(s []string, val string) {
	// set LAST element to val
	// then print s inside this function
	s[len(s)-1] = val
	fmt.Println(s)
}

func GrowSlice(s []string, val string) {
	// append val to s
	// then print s inside this function
	s = append(s, val)
	fmt.Println(s)
}

func main() {
	s := []string{"a", "b", "c"}

	fmt.Println("Before UpdateSlice:", s)
	UpdateSlice(s, "X")
	fmt.Println("After UpdateSlice:", s)

	fmt.Println("Before GrowSlice:", s)
	GrowSlice(s, "Y")
	fmt.Println("After GrowSlice:", s)
}
