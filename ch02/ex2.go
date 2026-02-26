package main

import "fmt"

func main() {
	const value = 20

	var i int = value
	var f float64 = value
	fmt.Printf(" i =%d  (%T)\n", i, i)
	fmt.Printf(" f =%f  (%T)\n", f, f)
}
