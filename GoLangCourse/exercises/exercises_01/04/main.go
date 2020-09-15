package main

import "fmt"

func main() {
	x := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Println(x, b, c, d, e, f)
}
