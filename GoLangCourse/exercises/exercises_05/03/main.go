package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("I am,", p.first, p.last, "and I am,", p.age)
}

func main() {
	p1 := person{
		first: "Diego",
		last:  "Molina",
		age:   30,
	}

	p1.speak()
}
