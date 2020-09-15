package main

import "fmt"

func main() {
	m := map[string][]string{
		"james_bond": []string{"shaken", "martinis", "women"},
		"no_dr": []string{
			"being evil", "ice cream", "sunset",
		},
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
