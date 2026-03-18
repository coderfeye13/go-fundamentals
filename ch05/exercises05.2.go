package main

import (
	"fmt"
	"os"
)

// fileLen takes a filename and returns the total number of bytes and an error.
func fileLen(filename string) (int, error) {
	f, err := os.Open(filename)
	if err != nil {
		return 0, err
	}

	// This is a critical pattern in DevOps to prevent resource leaks.
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	info, err := f.Stat()
	if err != nil {
		return 0, err
	}

	// Cast the int64 size to int to match the function signature
	return int(info.Size()), nil
}

func main() {
	fileName := "ch05/exercises05.2.go"

	size, err := fileLen(fileName)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Success: The size of '%s' is %d bytes.\n", fileName, size)
}
