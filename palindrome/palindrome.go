// Ver si el numero x se lee de igual manera de izq a derecha y de derecha a izq
// si se lee igual devolver true, si no false

package main

import (
	"strconv"
)

func IsPalindrome(x int) (bool, error) {
	xReverse := ""
	str := strconv.Itoa(x)
	// fmt.Println(str)
	for j := len(str) - 1; j >= 0; j-- {
		xReverse += string(str[j])
	}
	// fmt.Println(xReverse)
	return str == xReverse, nil
}
