package core

import (
	"fmt"
	"strings"
)

type Employee struct {
	Name  string
	Email string
	Age   int
}

func Validate(e Employee) (isValid bool, err error) {
	if e.Name == "" {
		return false, fmt.Errorf("name cannot be empty")
	}

	if e.Age < 18 {
		return false, fmt.Errorf("age must be greater than 18")
	}

	if !strings.Contains(e.Email, "@") {
		return false, fmt.Errorf("email must contain @")
	}

	return true, nil
}
