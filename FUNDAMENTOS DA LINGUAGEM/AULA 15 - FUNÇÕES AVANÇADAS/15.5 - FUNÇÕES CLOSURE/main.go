package main

import "fmt"

func closure() func() {
	texto := "Dentro da closure"

	funcao := func() {
		println(texto)
	}
	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)
	funcaoNova := closure()
	funcaoNova()
}
