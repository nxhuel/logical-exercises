package twosum_test

import (
	"ejercicios/twoSum"
	"testing"
)

func TestTwoSumResult(t *testing.T) {
	resultado, err := twosum.TwoSum([]int{3, 2, 4}, 6)
	if err != nil {
		t.Errorf("Se esperaba que no hubiera error, pero se obtuvo %v", err)
	}
	expected := []int{1, 2}
	if !sliceEqual(resultado, expected) {
		t.Errorf("Se esperaba %v pero se obtuvo %v", expected, resultado)
	}
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestTwoSumNoResult(t *testing.T) {
	resultado, err := twosum.TwoSum([]int{1, 2, 3}, 10)
	if err != nil {
		t.Errorf("Se esperaba un error, pero no se obtuvo ninguno")
	}

	if err != twosum.ErrorWithoutResult {
		t.Errorf("Se esperaba el error %v, pero se obtuvo %v", twosum.ErrorWithoutResult, err)
	}

	if resultado != nil {
		t.Errorf("Se esperaba un resultado nulo, pero se obtuvo %v", resultado)
	}
}
