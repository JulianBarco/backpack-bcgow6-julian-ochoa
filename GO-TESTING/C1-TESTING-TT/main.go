package main

import "fmt"

func Order(numeros []int) int {

	numeroMayor := numeros[0]
	for _, numero := range numeros {
		if numero > numeroMayor {
			numeroMayor = numero
		}
	}
	fmt.Printf("El número mayor entre %v es %d ", numeros, numeroMayor)
	return numeroMayor

}

func main() {
	Order([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
