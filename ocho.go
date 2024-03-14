// En este ejemplo, debe validar si una cadena de entrada de usuario es alfanumérica. La cadena dada no es , por lo que no tiene que verificarlo.nil/null/NULL/None

// La cadena tiene las siguientes condiciones para ser alfanumérica:

// Al menos un carácter (no es válido)""
// Los caracteres permitidos son letras latinas mayúsculas / minúsculas y dígitos de a 09
// Sin espacios en blanco / guiones bajos

// Comprobar si la cadena pasada es alfanumerica
// Tiene que tenger minimo un caracter != ""
// Tiene que ser menor a 10 y despues letra cualquiera
// no puede tener espacios ni guiones // listo

// package main

// import (
// 	"fmt"
// )

// func alphanumeric(str string) bool {
// 	var booleanos bool
// 	for _, i := range str {
// 		if (i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') && (i != ' ') && (i != '_') {
// 			booleanos = true
// 		} else {
// 			booleanos = false
// 		}
// 	}
// 	return booleanos
// }

// func main() {
// 	fmt.Println(alphanumeric("a"))
// 	fmt.Println(alphanumeric("Mazinkaiser"))
// 	fmt.Println(alphanumeric("hello world_"))
// 	fmt.Println(alphanumeric("PassW0rd"))
// 	fmt.Println(alphanumeric("     "))
// 	fmt.Println(alphanumeric(""))
// 	fmt.Println(alphanumeric("\n\t\n"))
// 	fmt.Println(alphanumeric("ciao\n$$_"))
// 	fmt.Println(alphanumeric("__ * __"))
// 	fmt.Println(alphanumeric("&)))((("))
// 	fmt.Println(alphanumeric("43534h56jmTHHF3k"))

// 	// Expect(alphanumeric("a")).To(Equal(true))
// 		// Expect(alphanumeric("Mazinkaiser")).To(Equal(true))
// 		// Expect(alphanumeric("hello world_")).To(Equal(false))
// 		// Expect(alphanumeric("PassW0rd")).To(Equal(true))
// 		// Expect(alphanumeric("     ")).To(Equal(false))
// 		// Expect(alphanumeric("")).To(Equal(false))
// 		// Expect(alphanumeric("\n\t\n")).To(Equal(false))
// 		// Expect(alphanumeric("ciao\n$$_")).To(Equal(false))
// 		// Expect(alphanumeric("__ * __")).To(Equal(false))
// 		// Expect(alphanumeric("&)))(((")).To(Equal(false))
// 		// Expect(alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
// }
