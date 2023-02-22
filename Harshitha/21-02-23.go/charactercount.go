package main

import (
	"fmt"
	"strings"
)

func main() {
	string1 := "how are you bro? wassup?"
	fmt.Println(string1)
	fmt.Println("Enter the character to be counted in the above string")
	var character string
	fmt.Scan(&character)
	fmt.Println(strings.Count(string1, character)) // 4
}
