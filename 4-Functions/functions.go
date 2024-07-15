package main

import "fmt"

func sum(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func mathCalc(n1, n2 int8) (int8, int8) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {
	sum := sum(10, 15)
	fmt.Println(sum)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Função por variável")

	resultSum, _ := mathCalc(5, 7)
	fmt.Print(resultSum)

}
