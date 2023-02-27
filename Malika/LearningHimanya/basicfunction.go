package main

import "fmt"

const a string = "Chandan"

func name(str string) {
	fmt.Println("My name is", str)
}

var username string

func main() {
	fmt.Scan(&username)
	name(username)
	//	name("123")
	username = "chandan"
	fmt.Println("This is main function", username)
	name(username)

	if username == "chandan" {
		var a int = 123
		fmt.Println(a)
	}

	fmt.Println(a)
}
