package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {

	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 1 {
		return strs[0]
	}

	firstWord := strs[0]
	for i := 0; i < len(strs); i++ {
		firstWordLength := len(firstWord)
		tempLength := len(strs[i])
		length := min(firstWordLength, tempLength)
		j := 0
		for j < length {
			if firstWord[j] != strs[i][j] {
				break
			}
			j++
		}
		firstWord = firstWord[:j]
		if len(firstWord) == 0 {
			return ""
		}
	}
	return firstWord
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}
