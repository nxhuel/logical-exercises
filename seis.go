// Defina una función que tome un argumento entero y devuelva un valor lógico o dependiendo de si el entero es un número primo.truefalse

// Según Wikipedia, un número primo (o un número primo) es un número natural mayor que que no tiene divisores positivos que no sean y sí mismo.11

// Requisitos
// Puede suponer que se le dará una entrada entera.
// No se puede suponer que el entero será solo positivo. Es posible que también se le den números negativos ( o ).0
// NOTA sobre el rendimiento: No se requieren optimizaciones sofisticadas, pero aún así las soluciones más triviales pueden agotar el tiempo de espera. Los números suben a 2^31 (o similar, dependiendo del idioma). El bucle hasta , o , será demasiado lento.nn/2

// Ejemplo
// is_prime(1)  /* false */
// is_prime(2)  /* true  */
// is_prime(-1) /* false */

// devolver true si el numero entero q se haya pasado es num primo
// los num no primos son negativos, 0 y 1
// los num primos son divisible todos numeros impares excpeto el 2

// package main

// import "fmt"

// func IsPrime(n int) bool {

// 	if n <= 1 {
// 		return false
// 	}
// 	for i := 2; i * i <= n; i++ {
// 		if n % i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func main() {
// 	fmt.Println(IsPrime(4))
// }
