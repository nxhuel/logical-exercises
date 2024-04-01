// Dado numeros enteros en una matriz y un numero entero en target, devolver los indices de los numeros que den el resultado en suma de target
// Ejemplo:
// Input: nums = [3,2,4], target = 6
// Output: [1,2]

// Llamar a la func q reciba matriz y target(int)
// Probar recorrer la lista de nums
// En la primera iteraccion q recorra el 1er indice y este aumente 1 en uno
// En la segunda iteraccion (q recorra el 2do indice y este aumente 1 en uno
// Primera y Segunda iteraccion se suma

// Errors:
// Se rompe cuando J supera el limite (es decir, I y J se chocan)
// Si no tiene solucion envia una matriz vacia

package twosum

func TwoSum(nums []int, target int) ([]int, error) {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				return result, nil
			}
		}
	}
	return nil, ErrorWithoutResult
}
