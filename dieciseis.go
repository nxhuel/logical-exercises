// Dado un número entero , return if es un xtruex palíndromo y de otro modo.false.
// Un entero es un palíndromo cuando se lee igual hacia adelante y hacia atrás. Por ej: 121
// Dado el num entero (x) iterar, probar que funcione solo iterar los numeros.
// Preguntar si el 1er numero y el ultimo coinciden, si es asi, true, si no, false. Para ubicar puedo convertir el num en string para poder separarlo
// Como puedo ubicar el ultimo numero? Quizas sabiendo la longitud, iteras sobre la longitud

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func isPalindrome(x int) bool {

// 	str := strconv.Itoa(x)

// 	firstDigit, _ := strconv.Atoi(string(str[0]))
// 	generalDigit, _ := strconv.Atoi(string(str))
// 	lastDigit, _ := strconv.Atoi(string(str[len(str)-1]))

// 	fmt.Println("firstDigit:", firstDigit)
// 	fmt.Println("generalDigit:", generalDigit)
// 	fmt.Println("lastDigit:", lastDigit)


// 	if generalDigit < 0 || firstDigit != lastDigit {
// 		return false
// 	}
// 	return true
// }

// func main() {
// 	// fmt.Println(isPalindrome(121))
// 	// fmt.Println(isPalindrome(-121))
// 	// fmt.Println(isPalindrome(10))
// 	// fmt.Println(isPalindrome(-10))
// 	// fmt.Println(isPalindrome(0))

// 	fmt.Println(isPalindrome(1000021))

// }
