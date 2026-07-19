package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoEstaAprovado(nota1, nota2 float32) bool {
	defer fmt.Println("Média calculada. Resultado da aprovação será retornado!")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado.")
	nota := (nota1 + nota2) / 2
	if nota >= 7 {
		return true
	} else {
		return false
	}
}

func main() {
	defer funcao1()
	// DEFER = Adiar a execução de uma função até o final da função que a chamou
	funcao2()
	fmt.Println(alunoEstaAprovado(7, 8))
}
