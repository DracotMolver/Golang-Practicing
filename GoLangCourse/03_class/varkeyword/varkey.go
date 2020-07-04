package main

import "fmt"

// It can be declared from outside a function
var x = 50

// You can assign the typ
var z int

func main() {
	var y = 43

	fmt.Println(y)

	z = 40
	fmt.Println(z)

	foo()
}

func foo() {
	fmt.Println(x)
}
