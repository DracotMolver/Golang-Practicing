package main

import "fmt"

func main() {
	result := f2c(49)

	fmt.Println(result)
}

func f2c(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}
