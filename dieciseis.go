// Retornar TRUE si es palindrome (numero que cuando se lee de adelante hacia atras y viceverse es igual)
// Crear una variable para el primero numero, y otra variable en el que se reserva para el numero ya invertido (averiguar como invertir un numero). Una vez el numero invertido comparo si ambas variables son iguales, si es asi es porques es palindrome (TRUE).
// Como invertir un numero? Con un ciclo for itero desde el ultimo indice de la longitud (len - 1) y va iterar mientras el indice sea mayor o igual a cero, este va a iterar restando -1 en la posicion.

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	str := strconv.Itoa(x)
	// fmt.Println(str)

	numInvertido := ""

	for i := len(str) -1; i >= 0; i-- {
		numInvertido += string(str[i])
	}
	// fmt.Println(numInvertido)
	if str == numInvertido{
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
	fmt.Println(isPalindrome(-10))
	fmt.Println(isPalindrome(0))
	fmt.Println(isPalindrome(1000021))
}
