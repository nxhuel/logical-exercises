// La raíz digital es la suma recursiva de todos los dígitos de un número.

// Dado , toma la suma de los dígitos de . Si ese valor tiene más de un dígito, continúe reduciendo de esta manera hasta que se produzca un número de un solo dígito. La entrada será un número entero no negativo.nn

// Ejemplos
// 16 --> 1 + 6 = 7
// 942 --> 9 + 4 + 2 = 15 --> 1 + 5 = 6
// 132189 --> 1 + 3 + 2 + 1 + 8 + 9 = 24 --> 2 + 4 = 6
// 493193 --> 4 + 9 + 3 + 1 + 9 + 3 = 29 --> 2 + 9 = 11 --> 1 + 1 = 2

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func DigitalRoot(n int) int {
// 	var resultado int
// 	str := strconv.Itoa(n)

// 	for _, digito := range str {
// 		digitoValor, _ := strconv.Atoi(string(digito))
// 		resultado += digitoValor
// 	}
// 	if resultado >= 10 {
// 		return DigitalRoot(resultado)
// 	}

// 	return resultado
// }

// func main() {
// 	// fmt.Println(DigitalRoot(942))
// 	fmt.Println(DigitalRoot(2))
// }

