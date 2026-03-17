package main

import (
	"fmt"
)

// The inner function "captures" the prefix variable from the outer scope.
func prefixer(prefix string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s %s", prefix, name)
	}
}

func main() {
	// We initialize the closure with a specific prefix
	helloPrefix := prefixer("Hello")

	fmt.Println(helloPrefix("Furky"))
	fmt.Println(helloPrefix("Ewald"))

	logPrefix := prefixer("[LOG-INFO]")
	fmt.Println(logPrefix("Server started successfully"))
}
