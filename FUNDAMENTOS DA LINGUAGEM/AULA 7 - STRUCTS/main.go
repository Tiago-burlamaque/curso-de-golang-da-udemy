package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "tiago"
	u.idade = 17

	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos bobos", 140}

	usuario2 := usuario{"Tiago", 17, enderecoExemplo}

	fmt.Println(usuario2)
}
