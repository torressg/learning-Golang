package main

import "fmt"

func main() {
	var variavel string = "Variavel"
	variavel2 := "Variavel 2"
	fmt.Println(variavel)
	fmt.Println(variavel2)

	const (
		constante3 string = "Constante 1"
		constante4 int    = 2
	)
	fmt.Println(constante3, constante4)

	variavel5, variavel6 := 5, "Variavel 6"
	fmt.Println(variavel5, variavel6)

	variavel2, variavel5 = variavel6, variavel5

}
