// Escriba un analizador simple que analice y ejecute Deadfish.

// Deadfish tiene 4 comandos, cada uno de 1 carácter:

// i incrementa el valor (inicialmente 0)
// d disminuye el valor
// s eleva al cuadrado el valor
// o Imprime el valor en la matriz de retorno
// Los caracteres no válidos deben ignorarse.

// Parse("iiisdoso") == []int{8, 64}

// para realizar la busqueda de letras en la palabra puedo usar range, luego con indice esta en i pido que incremente un valor (contador), si llega a aparecer una d, se le resta uno, si aparece una s, los valores antes de la s se eleva todo al cuadrado, la o debo almacenar todo en una matriz/array

// package main

// import "fmt"

// func Parse(data string) []int {
// 	contador := 0

// 	var array []int

// 	for _, indice := range data {
// 		switch indice {
// 		case 'i':
// 			contador += 1
// 		case 'd':
// 			contador -= 1
// 		case 's':
// 			contador *= contador
// 		case 'o':
// 			array = append(array, contador)
// 		}
// 	}

// 	return array
// }

// func main() {
// 	fmt.Println(Parse("isoisoiso"))
// }
