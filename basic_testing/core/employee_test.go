package core_test

import (
	"ejercicios/basic_testing/core/core"
	"testing"
)

func TestValidate_EmptyName(t *testing.T) {
	e := core.Employee{
		Name: "",
	}

	isValid, err := core.Validate(e)

	if isValid != false {
		t.Error("expected isValid to be false")
	}
	if err == nil {
		t.Error("expected err to be not nill")
	}
	if err.Error() != "name cannot be empty" {
		t.Errorf("expected err.Error() to be 'name cannot be empty', got '%s'", err.Error())
	}
}

func TestValidate_IllegalAge(t *testing.T) {
	e := core.Employee{
		Name:  "Jairo",
		Age:   14,
		Email: "valid@gmail.com",
	}

	isValid, err := core.Validate(e)

	if isValid != false {
		t.Error("expected isValid to be false")
	}
	if err == nil {
		t.Error("expected err to be not nill")
	}

	if err.Error() != "age must be greater than 18" {
		t.Errorf("expected err.Error() to be 'age must be greater than 18', got '%s'", err.Error())
	}
}

func TestValidate_WrongEmail(t *testing.T) {
	e := core.Employee{
		Name:  "Jairo",
		Age:   18,
		Email: "invalidgmail.com",
	}

	isValid, err := core.Validate(e)

	if isValid != false {
		t.Error("expected isValid to be false")
	}
	if err == nil {
		t.Error("expected err to be not nill")
	}

	if err.Error() != "email must contain @" {
		t.Errorf("expected err.Error() to be 'email must contain @', got '%s'", err.Error())
	}
}

func TestValidate_ValidEmployee(t *testing.T) {
	e := core.Employee{
		Name:  "Jairo",
		Age:   18,
		Email: "valid@gmail.com",
	}

	isValid, err := core.Validate(e)

	if isValid != true {
		t.Error("expected isValid to be true")
	}
	if err != nil {
		t.Error("expected nill error")
	}
}
