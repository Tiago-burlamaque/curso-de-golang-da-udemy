package main

import "fmt"

type usuario struct {
	nome  string
	idade int
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuário %s no banco de dados...\n", u.nome)
}

func (u usuario) maiorIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
	fmt.Printf("Parabéns %s! Agora você tem %d anos.\n", u.nome, u.idade)
}

func main() {
	usuario1 := usuario{
		"João",
		30,
	}
	fmt.Println(usuario1)
	usuario1.salvar()
	usuario1.maiorIdade()

	usuario2 := usuario{
		"Tiago",
		17,
	}
	fmt.Println(usuario2)
	usuario2.salvar()
	usuario2.maiorIdade()
	fmt.Println(usuario1.maiorIdade())
	fmt.Println(usuario2.maiorIdade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}