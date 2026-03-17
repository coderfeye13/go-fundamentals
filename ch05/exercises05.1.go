package main

import (
	"errors"
	"fmt"
	"strconv"
)

type opFuncEx func(int, int) (int, error)

func addEx(i int, j int) (int, error) {
	return i + j, nil
}
func subEx(i int, j int) (int, error) {
	return i - j, nil
}
func mulEx(i int, j int) (int, error) {
	return i * j, nil
}
func divEx(i int, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

func main() {
	var opMap = map[string]opFuncEx{
		"+": addEx,
		"-": subEx,
		"*": mulEx,
		"/": divEx,
	}
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("invalid expression:", expression)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		op := expression[1]
		opFuncEx, ok := opMap[op]
		if !ok {
			fmt.Println("unsupported operator:", op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result, err := opFuncEx(p1, p2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(result)
	}
}
