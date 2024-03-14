// Convertir los valores decimal de r, g y b en hexadecimal
// El rango permitido es de 0 - 255, si el numero se pasa de ese rango se redondea al 255
// Decimal | Hexadecimal
// ——–|————
// 10 | A
// 11 | B
// 12 | C
// 13 | D
// 14 | E
// 15 | F

//	1
// 	1 : 16 = 1
//			 0
// 1 - 0 * 16

// package main

// import (
// 	"fmt"
// 	"strconv"
// 	"strings"
// )

// func RGB(r, g, b int) string {

// 	var resultado []string
// 	var residuo int

// 	rgb := []int{r, g, b}

// 	for _, index := range rgb {

// 		cociente := index / 16

// 		for cociente > 0 {
// 			cociente = cociente / 16
// 		}

// 		residuo = index - cociente*16
// 		for residuo > 9 {
// 			switch residuo {
// 			case 10:
// 				resultado = append(resultado, "A")
// 			case 11:
// 				resultado = append(resultado, "B")
// 			case 12:
// 				resultado = append(resultado, "C")
// 			case 13:
// 				resultado = append(resultado, "D")
// 			case 14:
// 				resultado = append(resultado, "E")
// 			case 15:
// 				resultado = append(resultado, "F")
// 			}
// 		}

// 		residuoStr := strconv.Itoa(residuo)

// 		if len(residuoStr) < 2 {
// 			residuoStr = "0" + residuoStr
// 		}
// 		resultado = append(resultado, residuoStr)
// 	}
// 	// fmt.Println(resultado)
// 	residuoStr := strings.Join(resultado, "")
// 	return residuoStr
// }

// func main() {
// 	fmt.Println(RGB(1, 2, 3))

// }
