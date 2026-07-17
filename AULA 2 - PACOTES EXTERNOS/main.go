package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}
