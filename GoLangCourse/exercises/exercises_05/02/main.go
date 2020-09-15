package main

import "fmt"

func main() {
	ve := []int{1, 2, 3, 4, 5, 6}

	x := foo(ve...)

	fmt.Println(x)
}

func foo(val ...int) int {
	sum := 0

	for _, value := range val {
		sum += value
	}

	return sum
}

func bar(val []int) int {
	return foo(val...)
}
