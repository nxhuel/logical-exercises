package main

import (
	"testing"
)

func TestPalindromeResult(t *testing.T) {
	result, _:= IsPalindrome(121)
	expected := true
	if result != expected {
		t.Errorf("expected %v, but obtained %v for number 121", expected, result)
	}
}

func TestPalindomeNoResult(t *testing.T) {
	result, err := IsPalindrome(123)
	if err != nil {
		t.Errorf("was expected some error, but none was obtained")
	}

	if result {
		t.Errorf("expected false for non-palindromic number 123, but obtained %v", result)
	}
}