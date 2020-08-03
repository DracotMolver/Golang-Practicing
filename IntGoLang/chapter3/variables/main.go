package main

import "fmt"

// Multiple variables
var (
	a = 5
	b = 10
	c = 15
)

const (
	d = 20
	e = 25
	f = 30
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)

	fmt.Println("Enter a number:")
	var input float64

	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)
}
