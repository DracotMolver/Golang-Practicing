package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"First"`
	Age   int    `json:"Age"`
}

func main() {
	data := `[{"First":"D","Age":31},{"First":"m","Age":45}]`

	var people []user

	err := json.Unmarshal([]byte(data), &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)
}
