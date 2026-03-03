package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	blocksAndShadowing()
	ifScopedInit()
	forExamples()
	rangeExamples()
	labelExample()
	switchExamples()
	gotoExample()
}

// Blocks & Shadowing
func blocksAndShadowing() {
	x := 10

	if x > 5 {
		x := 5 // shadows outer x (different variable)
		fmt.Println("inner x:", x)
	}

	fmt.Println("outer x:", x) // still 10
}

// if with scoped variable
func ifScopedInit() {
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// n exists only inside this if/else chain
	if n := rand.Intn(10); n > 5 {
		fmt.Println("big:", n)
	} else {
		fmt.Println("small:", n)
	}

	// fmt.Println(n) // compile error (out of scope)
}

// for styles
func forExamples() {

	// Classic C-style
	for i := 0; i < 3; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	// While-style
	i := 1
	for i < 10 {
		fmt.Print(i, " ")
		i *= 2
	}
	fmt.Println()

	// Infinite loop + break
	j := 0
	for {
		if j == 3 {
			break
		}
		fmt.Print(j, " ")
		j++
	}
	fmt.Println()
}

// for-range
func rangeExamples() {

	s := []int{2, 4, 6}

	// value is a COPY (does not modify slice)
	for _, v := range s {
		v *= 2
	}
	fmt.Println("slice unchanged:", s)

	// map iteration order is NOT guaranteed
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range m {
		fmt.Println(k, v)
	}

	// string range iterates over RUNES (not bytes)
	str := "apple_π!"
	for i, r := range str {
		fmt.Println(i, r, string(r))
	}
}

// Labels
func labelExample() {
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue outer // jump to next outer iteration
			}
			fmt.Println("i:", i, "j:", j)
		}
	}
}

// switch
func switchExamples() {

	n := 5

	switch n {
	case 1, 2:
		fmt.Println("small")
	case 5:
		fmt.Println("exactly five")
	default:
		fmt.Println("other")
	}

	// blank switch (like clean if/else chain)
	switch {
	case n < 3:
		fmt.Println("less than 3")
	case n > 3:
		fmt.Println("greater than 3")
	default:
		fmt.Println("equal 3")
	}
}

// goto
func gotoExample() {
	a := rand.Intn(10)

	for a < 30 {
		if a%5 == 0 {
			goto done // jump to shared exit logic
		}
		a = a*2 + 1
	}

done:
	fmt.Println("final a:", a)
}
