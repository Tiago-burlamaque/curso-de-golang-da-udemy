package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	fmt.Println(texto)
	for _, numero := range numeros {
		fmt.Println(numero)
	}
}

func main() {
	totalSoma := soma(2, 3, 4, 5, 5, 1, 6, 1241, 54)
	fmt.Println(totalSoma)

	escrever("Números informados:", 2, 3, 4, 5, 5, 1, 6, 1241, 54)
}
