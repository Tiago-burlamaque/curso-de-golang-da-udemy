package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroinvertido := inverterSinal(numero)
	fmt.Println(numeroinvertido)

	novoNumero := 40
	fmt.Println("Antes da função:", novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println("Depois da função:", novoNumero)
}
