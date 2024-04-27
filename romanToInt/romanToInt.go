// De romano a entero
package main

import (
	"fmt"
)

func invertirS(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func romanToInt(s string) int {
	romanInt := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	resultado := 0

	s = invertirS(s)

	for _, letra := range s {
		numero := romanInt[letra]
		if numero*3 > resultado {
			resultado += numero
		} else {
			resultado -= numero
		}
	}
	return resultado
}

func main() {
	fmt.Println(romanToInt("XV"))
}
