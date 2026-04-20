package main

import (
	"fmt"
	"time"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	start := time.Now()

	// Part 1 — without capacity
	//student := []Student{}
	// Part 2
	student := make([]Student, 0, 10_000_000)

	for i := 0; i < 10_000_000; i++ {
		student = append(student, Student{
			FirstName: "Ali",
			LastName:  "Yılmaz",
			Age:       30,
		})
	}

	fmt.Printf("Done! len=%d time=%s\n", len(student), time.Since(start))
}
