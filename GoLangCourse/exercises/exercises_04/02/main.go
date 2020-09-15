package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 5, 6}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
