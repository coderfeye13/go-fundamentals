package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Exercise 1: create a slice and fill it with 100 random numbers (0–100)
	var numbers []int

	for i := 0; i < 100; i++ {
		random := rand.Intn(101)
		numbers = append(numbers, random)
	}

	// Exercise 2: solution using if / else
	for _, n := range numbers {

		if n%6 == 0 {
			fmt.Println("Six!", n)
		} else if n%2 == 0 {
			fmt.Println("Two!", n)
		} else if n%3 == 0 {
			fmt.Println("Three!", n)
		} else {
			fmt.Println("Never mind!", n)
		}
	}

	// Exercise 2: solution using switch
	for _, n := range numbers {

		switch {
		case n%6 == 0:
			fmt.Println("Six!", n)
		case n%2 == 0:
			fmt.Println("Two!", n)
		case n%3 == 0:
			fmt.Println("Three!", n)
		default:
			fmt.Println("Never mind!", n)
		}

	}
}
