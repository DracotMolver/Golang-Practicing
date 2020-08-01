package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := person{
		firstname: "diego",
		lastname:  "molina",
		age:       32,
	}

	fmt.Println(p1, p1.firstname)

	p2 := secretAgent{
		person: person{
			firstname: "diego killer",
			lastname:  "molina",
		},
		ltk: true,
	}

	fmt.Println(p2, p2.lastname)

	// anonymous struct
	p3 := struct {
		firstname string
		inlove    bool
	}{
		firstname: "drake",
		inlove:    true,
	}

	fmt.Println(p3)
}
