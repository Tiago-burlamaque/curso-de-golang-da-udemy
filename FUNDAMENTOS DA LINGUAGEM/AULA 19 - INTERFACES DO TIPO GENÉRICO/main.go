package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("Olá, mundo!")
	generica(42)
	generica(3.14)

	mapa := map[interface{}]interface{}{
		1: "Olá, mundo!",
		2: 42,
		3: 3.14,
	}
	fmt.Println(mapa)
}
