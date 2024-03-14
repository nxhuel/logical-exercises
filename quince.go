// Dado los numeros del array nums, iterar sobre el array nums y almarcenar dos indices que sumer el resultado de target y almacenarlo en otro array X
// En el array X solo debe tener dos indices, sumar ambos numeros, el resultado debe ser igual a target
// Sumar 0 1, 0, 2, 0 3
// Iterar para cada numero, es decir, 0 va a sumar con 0 0, 0, 1, 0 2, 0 3 hasta q acabe la long. Despues va a empezar el 1 0, 1 1, 1 2, 1 3, y asi sucesivamente

// package main

// import "fmt"

// func twoSum(nums []int, target int) []int {

// 	newNums := []int{}

// 	for i := 0; i < len(nums); i++ {
// 		for j := 0; j < len(nums); j ++ {
// 			if i != j {
// 				sum := nums[i] + nums[j]
// 				if sum == target {
// 					newNums = append(newNums, i)
// 					newNums = append(newNums, j)
// 					return newNums
// 				}
// 			}
// 		}
// 	}
// 	return newNums
// }

// func main() {
// 	result := twoSum([]int{3, 2, 4}, 6)
// 	fmt.Println(result)
// }
