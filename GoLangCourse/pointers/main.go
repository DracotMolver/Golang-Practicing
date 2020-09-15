package main

import "fmt"

func main() {
	a := "diego"

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	var b *string = &a

	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)

	x := 0

	foo(&x)
	fmt.Println(x)
}

func foo(y *int) {
	fmt.Println(y)
	*y = 43
	fmt.Println(y)
}
