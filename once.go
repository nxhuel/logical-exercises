// Si tuviéramos que montar un juego de tres en raya, querríamos saber si el estado actual del tablero está resuelto, ¿no es así? ¡Nuestro objetivo es crear una función que lo compruebe por nosotros!

// Supongamos que el tablero viene en forma de una matriz de 3x3, donde el valor es si un lugar está vacío, si es una "X" o si es una "O", así:012

// [[0, 0, 1],
//  [0, 1, 2],
//  [2, 1, 0]]
// Queremos que nuestra función devuelva:

// -1 si el tablero aún no está terminado Y nadie ha ganado aún (hay espacios vacíos),
// 1 si "X" ganaba,
// 2 si ganaba "O",
// 0 si es un juego de gatos (es decir, un empate).

// package main

// import (
// 	"fmt"
// )

// func IsSolved(board [3][3]int) int {
// 	for i := 0; i < 3; i++ {
// 		if board[i][0] != 0 && board[i][0] == board[i][1] && board[i][0] == board[i][2] {
// 			return board[i][0]
// 		}
// 		if board[0][i] != 0 && board[0][i] == board[1][i] && board[0][i] == board[2][i] {
// 			return board[0][i]
// 		}
// 	}

// 	if board[0][0] != 0 && board[0][0] == board[1][1] && board[0][0] == board[2][2] {
// 		return board[0][0]
// 	}
// 	if board[0][2] != 0 && board[0][2] == board[1][1] && board[0][2] == board[2][0] {
// 		return board[0][2]
// 	}

// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			if board[i][j] == 0 {
// 				return -1 
// 			}
// 		}
// 	}

// 	return 0
// }

// func main() {
// 	board := [3][3]int{
// 		{1, 1, 1},
// 		{0, 2, 2},
// 		{0, 0, 0},
// 	}
// 	fmt.Println(IsSolved(board))
// }
