package main

import "fmt"

func main() {
	const (
		a = 2020
		b = 2020 - iota
		c = 2020 - iota
		d = 2020 - iota
	)

	fmt.Println(a, b, c, d)
}
