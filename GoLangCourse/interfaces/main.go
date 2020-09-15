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

type human interface {
	speak()
}

func foo() {
	fmt.Println("foo")
}

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I called person", h.(person).first)
	case secretAgent:
		fmt.Println("I called person", h.(secretAgent).first)

	}

	fmt.Println("I called human")
}

func (s secretAgent) speak() {
	fmt.Println("I am a super agent", s.first, s.last)
}

func (s person) speak() {
	fmt.Println("I am a person", s.first, s.last)
}

func main() {
	defer foo()

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

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println("----------------------")
	bar(sa1)
	bar(sa2)
	bar(p1)
}
