// Persona recibe un bono -> $???
// El bono se lo gasta en cervezas (planea hacer piramides con esas)
// Mientras el bono sera mayor al precio de la cerveza, este siga comprando

// Piramide de latas: cada nivel se eleva al cuadrado. Si estan en el nivel 1, 1, nivel 2, 4, nivel 3, 9, nivel 4, 16, etc
// Devolver en la func el numero de niveles completos

// For example: beeramid(1500, 2); // should === 12
// 2 x 750 = 1500

// package main

// import "fmt"

// func Beeramid(bonus int, price float64) int {

// 	var contador int
// 	longitud := []int{}

// 	for i := int(price); i < bonus; i++ {
// 		i = i * i
// 		longitud = append(longitud, i)
// 		for j := 0; j <= len(longitud); j++ {
// 			contador = j
// 		}
// 	}
// 	// fmt.Println(longitud)
// 	return contador
// }

// func main() {
// 	fmt.Println(Beeramid(9, 2.0))
// }
