// Se te va a dar una palabra. Tu trabajo es devolver el carácter medio de la palabra. Si la longitud de la palabra es impar, devuelve el carácter medio. Si la longitud de la palabra es par, devuelve los 2 caracteres del medio.

// 1. Pasar una palabra por parametro
// 2. Devolver de la palabra, el caracter el medio
// 2.1. con un for contar letra por letra, verificar que se separen para luego validar si es par o impar
// 3. Si es par, devuelvo las dos letras del medio, si es impar, solo devuelvo la letra del medio

// package main

// import (
// 	"fmt"
// )

// func GetMiddle(s string) string {
	
// 	var letras string
// 	for i := 0; i < len(s); i++ {
// 		letra := string(s[i])
// 		letras += letra
// 	}
	
// 	var resultado string
// 	if len(letras) % 2 == 0 {
// 		middleStart := len(letras)/2 - 1 
// 		middleEnd := len(letras)/2 + 1    
// 		for j := middleStart; j < middleEnd; j++ {
// 			resultado += string(letras[j])
// 		}
// 	} else {
// 		middleIndex := len(letras) / 2
// 		resultado = string(letras[middleIndex])
// 	}
// 	return resultado
// }

// func main() {
// 	g := GetMiddle("testing")
// 	fmt.Printf(g)
}