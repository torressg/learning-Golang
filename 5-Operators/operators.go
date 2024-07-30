package main

import "fmt"

func main() {
	sum := 1 + 2
	sub := 1 - 2
	div := 10 / 4
	mult := 10 * 5
	remainder := 12 % 6

	fmt.Println(sum, sub, div, mult, remainder)

	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(!verdadeiro)

	numero := 10
	numero++
	numero += 15 // numero = numero + 15
}
