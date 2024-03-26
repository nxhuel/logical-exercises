package main

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{-121, false},
		{10, false},
		{-10, false},
		{0, true},
		{1000021, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Input: %d", tc.input), func(t *testing.T) {
			result := isPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("Expected isPalindrome(%d) to be %t, but got %t", tc.input, tc.expected, result)
			}
		})
	}
}
