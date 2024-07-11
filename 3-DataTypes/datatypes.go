package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 100
	fmt.Println(numero)

	// INT32 = RUNE
	var numero2 rune = 123
	fmt.Println(numero2)

	// INT8 = BYTE
	var numero3 byte = 123
	fmt.Println(numero3)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 123.45
	fmt.Println(numeroReal2)

	var erro error = errors.New("Error found")
	fmt.Println(erro)

}
