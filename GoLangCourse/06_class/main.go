package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Println("hello, playground")
	}

	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop:%d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}

	x := 1
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	z := 1

	for {
		if z > 9 {
			break
		}

		fmt.Println(z)
		z++
	}
	fmt.Println("done.")

	a := 0

	for {
		a++
		if a > 100 {
			break
		}

		if a%2 != 0 {
			continue
		}

		fmt.Println(a)
	}
	fmt.Println("done.")
}
