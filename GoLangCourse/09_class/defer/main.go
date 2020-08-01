package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func main() {
	defer foo()
	bar()

	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"miss",
			"moneypenny",
		},
		ltk: true,
	}

	fmt.Println(sa1, sa2)
	sa1.speak()
	sa2.speak()
}
