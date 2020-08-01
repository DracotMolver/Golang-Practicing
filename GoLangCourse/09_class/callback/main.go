package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)

	fmt.Println(s)

	ee := even(sum, ii...)
	fmt.Println(ee)

}

func sum(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}

func even(fn func(x ...int) int, vi ...int) int {
	var yi []int

	for _, v := range vi {
		if v&1 == 0 {
			yi = append(yi, v)
		}
	}

	return fn(yi...)
}
