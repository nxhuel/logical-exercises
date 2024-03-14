// Dados dos arrays devolver true si son IGUALES o si los elementos del array b es el cuadrado del elemento del array a

//a = [121, 144, 19, 161, 19, 144, 19, 11]
//b = [121, 14641, 20736, 361, 25921, 361, 20736, 361]

// a = [121, 144, 19, 161, 19, 144, 19, 11]
// b = [11*11, 121*121, 144*144, 19*19, 161*161, 19*19, 144*144, 19*19]

// package main

// import (
// 	"fmt"
// )

// func Comp(array1 []int, array2 []int) bool {

// 	if len(array1) != len(array2) {
// 		return false
// 	}

// 	mapCuadrado := make(map[int]int)
// 	for _, elementoDos := range array2 {
// 		cuadrado := elementoDos
// 		mapCuadrado[cuadrado]++
// 	}

// 	for _, elemento := range array1 {
// 		cuadrado := elemento * elemento

// 		if mapCuadrado[cuadrado] > 0 {
// 			mapCuadrado[cuadrado]--
// 		} else {
// 			return false
// 		}
// 	}

// 	return true
// }

// func main() {
// 	array1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
// 	array2 := []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}

// 	fmt.Println(Comp(array1, array2))
// }
