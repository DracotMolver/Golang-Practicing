package main

import "fmt"

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n)
	fmt.Println(x, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 1989, "Big Brother is Watching"
}
