package main

import "fmt"

func main() {

	usuario := map[string]string {
		"Nome": "Tiago",
		"Sobrenome": "Burlamaque",
	}

	fmt.Println(usuario["Nome"])
}