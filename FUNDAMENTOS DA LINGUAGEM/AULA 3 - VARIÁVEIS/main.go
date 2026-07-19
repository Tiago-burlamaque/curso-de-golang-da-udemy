package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1" // 1° jeito de salvar variáveis
	fmt.Println(variavel1)

	varivael2 := "Variavel 2" // 2° jeito de salvar variavéis
	fmt.Println(varivael2)

	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	) // 3° jeito de salvar variavéis
	fmt.Println(variavel3)
	fmt.Println(variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6" // 4° jeito de salvar variaveis
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1" // Constante
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5 // Trocando variaveis
	fmt.Println(variavel5, variavel6)

}
