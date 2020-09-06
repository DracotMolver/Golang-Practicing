package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "John Doe",
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	// p.name = "Diego"
	(*p).name = "Diego"
}
