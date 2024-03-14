// Tome un número entero y un dígito como número entero. n (n >= 0)d (0 <= d <= 9)

// Eleva al cuadrado todos los números entre 0 y n. k (0 <= k <= n)

// Cuente el número de dígitos utilizados en la escritura de todos los archivos . dk**2

// Implemente la función tomando y como parámetros y devolviendo este recuento.nd

// Ejemplos:
// n = 10, d = 1
// the k*k are 0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100
// We are using the digit 1 in: 1, 16, 81, 100. The total count is then 4.

// The function, when given n = 25 and d = 1 as argument, should return 11 since
// the k*k that contain the digit 1 are:
// 1, 16, 81, 100, 121, 144, 169, 196, 361, 441.
// So there are 11 digits 1 for the squares of numbers between 0 and 25

// 1. Recibir un numero entero mayor igual a 0 (n) y un digito como entero que sea mayor igual a 0 y menor igual a 0 (d).
// 2. la variable k va a ser el cuadrado de todos los numeros entre 0 y n.
// 3. Almacenar los digitos en un array con variable (d) y esta va a contar el digito 1 de todos los numeros

// package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func NbDig(n, d int) {

// 	numeros := []int{}
// 	for i := 0; i <= n; i++ {
// 		numeros = append(numeros, i)
// 	}

// 	fmt.Println(numeros)

// 	numerosMultiplicados := []int{}
// 	for j := 0; j < len(numeros); j++ {
// 		numeroCuadrado := numeros[j] * numeros[j]
// 		numerosMultiplicados = append(numerosMultiplicados, numeroCuadrado)
// 	}

// 	fmt.Println(numerosMultiplicados)

// 	contarUnosEnNumero := func(numero int) int {
// 		numeroStr := strconv.Itoa(numero)
// 		contador := 0
// 		for _, digito := range numeroStr {
// 			if digito == '1' {
// 				contador++
// 			}
// 		}
// 		return contador
// 	}

// 	contadorDeUnos := 0
// 	for _, numero := range numerosMultiplicados {
// 		contadorDeUnos += contarUnosEnNumero(numero)
// 	}

// 	fmt.Println(contadorDeUnos)
// }

// func main() {
// 	NbDig(10, 1)
// }
