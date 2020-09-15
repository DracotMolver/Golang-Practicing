package main

import "fmt"

func main() {
	// short declaration belong to the scope of the function
	// can't be declared outside of a function body, for that
	// use `var`
	x := 42

	fmt.Println(x)

	y := 100 + 24

	fmt.Println(y)

	x = 99
	fmt.Println(x)
}
