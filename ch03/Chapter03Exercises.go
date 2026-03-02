package main

import "fmt"

func main() {

	greetings := []string{
		"Hello",
		"Hola",
		"Merhaba",
		"नमस्कार",
		"こんにちは",
		"Привіт",
	}

	firstTwo := greetings[:2]
	secondSet := greetings[1:5]
	thirdSet := greetings[3:6]

	fmt.Println("Original: ", greetings)
	fmt.Println("First Two: ", firstTwo)
	fmt.Println("Second set: ", secondSet)
	fmt.Println("Third set: ", thirdSet)

	message := "Hi and "

	runes := []rune(message)
	fmt.Println("Fourth rune: ", string(runes[3]))

	e1 := Employee{"Furkan", "Yilmaz", 1}

	e2 := Employee{
		FirstName: "Sila",
		LastName:  "Ozdemir",
		id:        2,
	}
	var e3 Employee
	e3.FirstName = "Sila"
	e3.LastName = "Yilmaz"
	e3.id = 3

	fmt.Println("Employee 1:", e1)
	fmt.Println("Employee 2:", e2)
	fmt.Println("Employee 3:", e3)

}

type Employee struct {
	FirstName string
	LastName  string
	id        int
}
