package main

import "fmt"

func main() {
	// composite literal - Values of the same type
	x := []int{4, 5, 6, 7, 8, 42}

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[0])

	for i, v := range x {
		fmt.Println(i, v)
	}

	// SLICING
	fmt.Println("Slicing")
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}

	// Append values
	x = append(x, 10, 67, 233)
	fmt.Println(x)

	y := []int{3, 64, 23, 75, 23}

	x = append(x, y...)
	fmt.Println(x)

	// Makes. work with an underline array. adding a new one with a double capacity when
	// the length is passed by
	a := make([]int, 10, 11)
	fmt.Println(a, len(a), cap(a))
	a = append(a, 12)
	a = append(a, 13)
	fmt.Println(a, len(a), cap(a))

	// multi-dimensional slices
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	jc := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(jc)

	xp := [][]string{jb, jc}
	fmt.Println(xp)
}
