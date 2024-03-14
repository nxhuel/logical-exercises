// Si enumeramos todos los números naturales por debajo de 10 que son múltiplos de 3 o 5, obtenemos 3, 5, 6 y 9. La suma de estos múltiplos es 23.

// Termina la solución de modo que devuelva la suma de todos los múltiplos de 3 o 5 por debajo del número pasado.

// Nota: Si el número es múltiplo de 3 y 5, cuéntelo solo una vez.

//	1. Recibir un numero entero que va a ser el limite de los multiplos de 3 o 5.
//	2. Puedo hacer un for para recorrer del 0 hasta el number y validar/retornar solo los multiplos de 3 o 5
//	3. Crear un array para guardar los multiplos de 3 o 5
// 	4. Sumar todo lo que hay dentro de multiplos

// package main

// import "fmt"

// func Multiple3And5(number int) int {

// 	var multiplos []int

// 	for i := 1; i < number; i++ {
// 		if i % 3 == 0 || i % 5 == 0 {
// 			multiplos = append(multiplos, i)
// 		}
// 	}

// 	var resultado int
// 	for _, suma := range multiplos {
// 		resultado += suma
// 	}
// 	return resultado
// }

// func main() {
// 	fmt.Println(Multiple3And5(10))
// }