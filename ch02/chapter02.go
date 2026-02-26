package main

import "fmt"

const untypedConst = 10
const typedConst int = 20

func main() {

	fmt.Println("===== PREDECLARED TYPES =====")

	var i int = 10
	var f float64 = 3.14
	var b bool = true
	var s string = "Go"

	fmt.Println(i, f, b, s)

	fmt.Println("\n===== ZERO VALUES =====")

	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	fmt.Println(zeroInt, zeroFloat, zeroBool, zeroString)

	fmt.Println("\n===== LITERALS =====")

	fmt.Println(100)     // int literal
	fmt.Println(0b1010)  // binary
	fmt.Println(0xFF)    // hex
	fmt.Println(3.14)    // float
	fmt.Println("hello") // string
	fmt.Println('A')     // rune
	fmt.Println(5i)      // complex literal

	fmt.Println("\n===== TYPE CONVERSION =====")

	var x int = 10
	var y float64 = 3.5

	sum1 := float64(x) + y
	sum2 := x + int(y)

	fmt.Println("sum1:", sum1)
	fmt.Println("sum2:", sum2)

	fmt.Println("\n===== CONSTANTS =====")

	fmt.Println("untypedConst:", untypedConst)
	fmt.Println("typedConst:", typedConst)

	var a float64 = untypedConst
	fmt.Println("untyped assigned to float:", a)

	// var c float64 = typedConst  // compile error

	fmt.Println("\n===== UNTYPED CONSTANT FLEXIBILITY =====")

	var intVar int = untypedConst
	var floatVar float64 = untypedConst
	var byteVar byte = untypedConst

	fmt.Println(intVar, floatVar, byteVar)

	fmt.Println("\n===== INVALID CONST EXAMPLE =====")

	// Uncomment to see compile error:
	// x := 5
	// y := 10
	// const z = x + y // runtime values cannot be const

}
