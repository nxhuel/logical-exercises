package main

import (
	"fmt"
	"ejercicios/basic_testing/core/core"
)

func main() {
	e := core.Employee{
		Name:  "Jairo",
		Email: "email@domain.com",
		Age:   19,
	}

	isValid, err := core.Validate(e)
	fmt.Println()
	fmt.Println("Is valid: ", isValid)
	fmt.Println("Error: ", err)
}
