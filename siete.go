// Escriba una función que tome un número entero como entrada y devuelva el número de bits que son iguales a uno en la representación binaria de ese número. Puede garantizar que la entrada no sea negativa.

// Ejemplo: La representación binaria de es , por lo que la función debe devolver en este caso1234100110100105

// package main

// import "fmt"

// func CountBits(numero uint) int {
// 	var resto []int
// 	for i := numero; i > 0; i /= 2 {
// 		resto = append(resto, int(i%2))
// 	}

// 	var resultado int
// 	for _, bit := range resto {
// 		resultado += bit
// 	}

// 	return resultado
// }

// func main() {
// 	numero := 10

// 	fmt.Printf("La cantidad de Unos en el codigo binario de %v es %v", numero, CountBits(uint(numero)))
// }
