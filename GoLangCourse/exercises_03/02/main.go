package main

import "fmt"

func main() {
	for i := 66; i < 91; i++ {
		fmt.Print(i, "\n")
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
