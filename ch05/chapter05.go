package main

import (
	"errors"
	"fmt"
	"sort"
)

// function type
type opFunc func(int, int) int

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	fmt.Println("Chapter 5 - Functions")

	basicFunctionDemo()
	namedOptionalStyleDemo()
	variadicDemo()
	multipleReturnDemo()
	namedReturnDemo()
	functionValueDemo()
	anonymousFunctionDemo()
	closureDemo()
	functionAsParameterDemo()
	returnFunctionDemo()
	deferDemo()
	callByValueDemo()
	mapSliceCallByValueDemo()
}

// basic function
func div(num, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

func basicFunctionDemo() {
	fmt.Println("\nBasic function:")
	result := div(5, 2)
	fmt.Println(result)
}

// struct as options
func myFunc(opts MyFuncOpts) error {
	fmt.Println(opts.FirstName, opts.LastName, opts.Age)
	return nil
}

func namedOptionalStyleDemo() {
	fmt.Println("\nStruct as options:")
	_ = myFunc(MyFuncOpts{
		LastName: "Patel",
		Age:      50,
	})
	_ = myFunc(MyFuncOpts{
		FirstName: "Joe",
		LastName:  "Smith",
	})
}

// variadic
func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func variadicDemo() {
	fmt.Println("\nVariadic:")
	fmt.Println(addTo(3))
	fmt.Println(addTo(3, 2))
	fmt.Println(addTo(3, 2, 4, 6, 8))

	a := []int{4, 3}
	fmt.Println(addTo(3, a...))
}

// multiple returns
func divAndRemainder(num, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("cannot divide by zero")
	}
	return num / denom, num % denom, nil
}

func multipleReturnDemo() {
	fmt.Println("\nMultiple returns:")
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result, remainder)

	// ignore one value
	result2, _, err2 := divAndRemainder(7, 3)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(result2)
}

// named returns
func divNamed(num, denom int) (result int, remainder int, err error) {
	if denom == 0 {
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func namedReturnDemo() {
	fmt.Println("\nNamed returns:")
	x, y, z := divNamed(10, 3)
	fmt.Println(x, y, z)
}

// functions are values
func f1(a string) int {
	return len(a)
}

func f2(a string) int {
	total := 0
	for _, r := range a {
		total += int(r)
	}
	return total
}

func functionValueDemo() {
	fmt.Println("\nFunctions are values:")

	var fn func(string) int

	fn = f1
	fmt.Println(fn("Hello"))

	fn = f2
	fmt.Println(fn("Hello"))

	ops := map[string]opFunc{
		"+": add,
		"-": sub,
		"*": mul,
		"/": divInt,
	}

	fmt.Println(ops["+"](2, 3))
	fmt.Println(ops["*"](2, 3))
}

func add(i, j int) int    { return i + j }
func sub(i, j int) int    { return i - j }
func mul(i, j int) int    { return i * j }
func divInt(i, j int) int { return i / j }

// anonymous function
func anonymousFunctionDemo() {
	fmt.Println("\nAnonymous functions:")

	f := func(j int) {
		fmt.Println("printing", j)
	}

	for i := 0; i < 3; i++ {
		f(i)
	}

	for i := 0; i < 3; i++ {
		func(j int) {
			fmt.Println("inline", j)
		}(i)
	}
}

// closure
func closureDemo() {
	fmt.Println("\nClosures:")

	a := 20
	f := func() {
		fmt.Println("inside before change:", a)
		a = 30
	}
	f()
	fmt.Println("outside after change:", a)

	b := 50
	g := func() {
		fmt.Println("before shadow:", b)
		b := 60
		fmt.Println("shadowed inner b:", b)
	}
	g()
	fmt.Println("outer b:", b)
}

// passing function as parameter
func functionAsParameterDemo() {
	fmt.Println("\nFunction as parameter:")

	people := []Person{
		{"Furkan", "Yilmaz", 25},
		{"Ewald", "Kanwald", 23},
		{"Juri", "Fredson", 18},
	}

	fmt.Println("original:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("by last name:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("by age:", people)
}

// returning function
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func returnFunctionDemo() {
	fmt.Println("\nReturning functions:")

	twoBase := makeMult(2)
	threeBase := makeMult(3)

	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}

// defer
func deferExample() int {
	a := 10

	// params evaluated immediately
	defer func(val int) {
		fmt.Println("first:", val)
	}(a)

	a = 20

	defer func(val int) {
		fmt.Println("second:", val)
	}(a)

	a = 30
	fmt.Println("exiting:", a)
	return a
}

func deferDemo() {
	fmt.Println("\nDefer:")
	result := deferExample()
	fmt.Println("returned:", result)
}

// call by value
type person struct {
	age  int
	name string
}

func modifyFails(i int, s string, p person) {
	i = i * 2
	s = "Goodbye"
	p.name = "Bob"
}

func callByValueDemo() {
	fmt.Println("\nCall by value:")

	p := person{}
	i := 2
	s := "Hello"

	modifyFails(i, s, p)
	fmt.Println(i, s, p)
}

// maps and slices
func modMap(m map[int]string) {
	m[2] = "hello"
	m[3] = "goodbye"
	delete(m, 1)
}

func modSlice(s []int) {
	for k, v := range s {
		s[k] = v * 2
	}
	s = append(s, 10)
}

func mapSliceCallByValueDemo() {
	fmt.Println("\nMaps and slices:")

	m := map[int]string{
		1: "first",
		2: "second",
	}
	modMap(m)
	fmt.Println(m)

	s := []int{1, 2, 3}
	modSlice(s)
	fmt.Println(s)
}
