// package main

// import (
// 	"fmt"
// 	"strconv" // importo strconv para convertir los str a int
// 	"strings" // importo strings para que en la var in no me lea los espacios
// )

// func HighAndLow(in string) string {

// 	max := 1 // Creo dos var, max y min de tipo int, ambos arrancan en uno con el proposito de que sean reemplazados por el max o min
// 	min := 1

// 	numerosStr := strings.Split(in, " ") // asigno la var numerosStr y uso strings.Split para que no lean los espacios

// 	for _, numStr := range numerosStr { // creo un for, con un indice que no se lee y un numStr, va iterando en el rango de numerosStr
// 		num, err := strconv.Atoi(numStr) // Hago uso del paquete strconv para convertir los strings dados a int
// 		if err != nil { // manejo el error 
// 			panic(err)
// 		}
// 		if num > max { // se itera numero x numero y si es mayor que el max de ese momento, se asigna al max, si no pasa de largo
// 			max = num
// 		}
// 		if num < min { // se itera numero x numero y si es menos que el min de ese momento, se asigna al min, si no pasa de largo
// 			min = num
// 		}
// 	}
// 	return strconv.Itoa(max) + " " + strconv.Itoa(min) // retorno ahora de int a str usando .Itoa los max y min
// }

// func main() {
// 	fmt.Println(HighAndLow("1 2 3"))
// }
