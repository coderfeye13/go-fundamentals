package main

import "fmt"

type Person struct {
	// Üç alan: FirstName, LastName string | Age int
	FirstName string
	LastName  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	// Person döndür (pointer değil)
	return Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	// *Person dondurur
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
	}
}

func main() {

	p1 := MakePerson("Ali", "Yılmaz", 25)
	p2 := MakePersonPointer("Ayşe", "Kaya", 30)

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Printf("%p\n", p2) //Real memory address

}
