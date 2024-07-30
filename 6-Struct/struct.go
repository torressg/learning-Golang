package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	street string
	number uint16
}

func main() {
	var user1 user
	user1.name = "Gui"
	user1.age = 21
	fmt.Println(user1)

	exampleAddress := address{"Silly St", 0}
	user2 := user{"Ma", 23, exampleAddress}
	fmt.Println(user2)

	user3 := user{age: 12}
	fmt.Println(user3)
}
