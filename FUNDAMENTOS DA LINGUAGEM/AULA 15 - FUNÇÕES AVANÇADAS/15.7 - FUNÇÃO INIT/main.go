package main

import "fmt"

var n int

func init() {
	fmt.Println("Função init está sendo executada.")
	n = 10
}

func main() {
	fmt.Println("Função main está sendo executada.")
	fmt.Println("Valor de n:", n)
}
