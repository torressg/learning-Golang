package main

import (
	"fmt"
	"module/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo da Main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("teste@gmail.com")
	fmt.Println((erro))
}
