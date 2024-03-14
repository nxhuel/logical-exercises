// Jaden Smith, el hijo de Will Smith, es el protagonista de películas como The Karate Kid (2010) y After Earth (2013). Jaden también es conocido por parte de su filosofía que transmite a través de Twitter. Cuando escribe en Twitter, es conocido por escribir casi siempre en mayúsculas cada palabra. Para simplificar, tendrás que poner en mayúsculas cada palabra, comprueba cómo se espera que sean las contracciones en el siguiente ejemplo.

// Su tarea es convertir las cadenas a la forma en que las escribiría Jaden Smith. Las cadenas son citas reales de Jaden Smith, pero no están en mayúsculas de la misma manera en que las escribió originalmente.

// Modificar la primera letra de cada palabra
// 1. Separar cada palabra
// 1. Una vez pasado las palabras con un for puedo recorrer del 0 hasta q termine la long cada palabra, primero probar que mi for funciona reccoriendo todas las palabras
// 2. Dentro del for hago uso del metodo .title que agarra la 1era letra de cada palabra que se separo y la convierte en mayuscula
// 3. Volver a unir todas las palabras en una oracion con espacios

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func ToJadenCase(str string) string {
// 	letras := strings.Fields(str)

// 	for i := 0; i < len(letras); i++ {
// 		letras[i] = strings.Title(letras[i])
// 	}

// 	resultado := strings.Join(letras, " ")
// 	return resultado
// }

// func main() {
// 	fmt.Println(ToJadenCase("most trees are blue"))
// }
