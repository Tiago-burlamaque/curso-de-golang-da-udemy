package main

import "fmt"

func somar(a int8, b int8) int8 {
	return a + b
}

func calculosMatematicos(a, b int8) (int8, int8) {
	soma := a + b
	subtracao := a - b
	fmt.Println("O resultado da soma é:", soma)
	fmt.Println("O resultado da subtração é:", subtracao)
	return soma, subtracao
}

func main() {
	soma := somar(6, 7)
	fmt.Println(soma)

	var f = func(txt string) string {
		fmt.Println("Função f")
		fmt.Println(txt)
		return txt
	}

	f("olá")
	calculosMatematicos(6, 7)
}
