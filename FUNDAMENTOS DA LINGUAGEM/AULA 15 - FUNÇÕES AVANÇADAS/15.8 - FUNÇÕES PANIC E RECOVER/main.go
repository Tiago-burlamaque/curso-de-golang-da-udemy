package main

import "fmt"

func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Recuperando da execução! Erro:", r)
	}
}

func alunoestaaprovado(nota1, nota2 float64) bool {
	defer recuperarExecucao()
	media := (nota1 + nota2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println(alunoestaaprovado(6, 6))
	fmt.Println("Pós execução!")
}
