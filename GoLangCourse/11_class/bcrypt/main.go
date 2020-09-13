package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"

	bc, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(s)
	fmt.Println(bc)

	loginPassword := "password123"

	err = bcrypt.CompareHashAndPassword(bc, []byte(loginPassword))

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("You are logged in")
}
