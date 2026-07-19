package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p := pessoa{"João", "pedro", 19, 190}
	fmt.Println(p)
	
	e := estudante{p, "Analise e desenvolvimento de sistemas", "IFSC"}
	fmt.Println(e.pessoa.altura)
}
