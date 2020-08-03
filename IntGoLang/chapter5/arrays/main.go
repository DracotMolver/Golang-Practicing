package main

import "fmt"

func main() {
	var x [5]int

	x[4] = 100

	fmt.Println(x)

	var z [5]float64

	// Shorten syntax
	y := [5]int{1, 2, 3, 4, 5}

	fmt.Println(y)

	z[0] = 98
	z[1] = 93
	z[2] = 77
	z[3] = 82
	z[4] = 83

	// Using direct number
	var total float64 = 0

	for i := 0; i < 5; i++ {
		total += z[i]
	}

	fmt.Println(total / 5)

	// Using the len function
	var total2 float64 = 0

	for i := 0; i < len(z); i++ {
		total2 += z[i]
	}

	fmt.Println(total2 / float64(len(z)))

	// Using the range keyword
	var total3 float64 = 0

	for _, v := range z {
		total3 += v
	}

	fmt.Println(total3 / float64(len(z)))
}
