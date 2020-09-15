package main

import "fmt"

func main() {
	xs1 := []string{"james", "bond", "shaken, no stirred"}
	xs2 := []string{"miss", "moneypenny", "hello, james."}

	x := [][]string{xs1, xs2}

	for _, v := range x {
		fmt.Println(v)
	}
}
