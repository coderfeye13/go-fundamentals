package main

import "fmt"

//Basic Struct Definition

type Person struct {
	name string
	age  int
	pet  string
}

//Struct for Comparison Demo

type FirstPerson struct {
	name string
	age  int
}

type SecondPerson struct {
	name string
	age  int
}

//Struct With Slice (Not Comparable)

type PersonWithHobbies struct {
	name    string
	hobbies []string
}

func main() {

	fmt.Println("===== ZERO VALUE STRUCT =====")
	var p1 Person
	fmt.Println(p1) // all fields zero value

	fmt.Println("\n===== STRUCT LITERAL (POSITIONAL) =====")
	p2 := Person{"Julia", 40, "cat"}
	fmt.Println(p2)

	fmt.Println("\n===== STRUCT LITERAL (NAMED FIELDS) =====")
	p3 := Person{
		age:  30,
		name: "Beth",
	}
	fmt.Println(p3)

	fmt.Println("\n===== FIELD ACCESS =====")
	p3.pet = "dog"
	fmt.Println("Name:", p3.name)
	fmt.Println("Pet:", p3.pet)

	fmt.Println("\n===== STRUCT COMPARISON =====")
	a := Person{"Alice", 25, "fish"}
	b := Person{"Alice", 25, "fish"}

	fmt.Println("a == b:", a == b) // true (all fields comparable)

	fmt.Println("\n===== STRUCT TYPE CONVERSION =====")
	fp := FirstPerson{"Bob", 50}
	sp := SecondPerson(fp) // allowed (same fields, same order)
	fmt.Println("Converted:", sp)

	fmt.Println("\n===== ANONYMOUS STRUCT =====")

	anon := struct {
		title string
		year  int
	}{
		title: "Go Fundamentals",
		year:  2026,
	}

	fmt.Println(anon)

	// Anonymous - Named comparison (same structure)
	var another struct {
		name string
		age  int
	}

	another = fp
	fmt.Println("Anonymous assignment works:", another)
	fmt.Println("Comparison works:", fp == another)

	fmt.Println("\n===== STRUCT WITH SLICE (NOT COMPARABLE) =====")

	ps1 := PersonWithHobbies{
		name:    "Charlie",
		hobbies: []string{"coding", "gym"},
	}

	ps2 := PersonWithHobbies{
		name:    "Charlie",
		hobbies: []string{"coding", "gym"},
	}

	fmt.Println(ps1)
	fmt.Println(ps2)

	// This would NOT compile:
	// fmt.Println(ps1 == ps2)

	fmt.Println("Struct with slice field cannot be compared using ==")

	fmt.Println("\n===== STRUCT IN SLICE =====")

	people := []Person{
		{"Tom", 20, "parrot"},
		{"Emma", 28, "cat"},
		{"Liam", 35, "dog"},
	}

	for _, person := range people {
		fmt.Printf("%s is %d years old and has a %s\n",
			person.name,
			person.age,
			person.pet,
		)
	}
}
