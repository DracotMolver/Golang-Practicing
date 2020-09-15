package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "John",
		Last:  "Doe",
		Age:   30,
	}

	p2 := person{
		First: "Diego",
		Last:  "Molina",
		Age:   31,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
