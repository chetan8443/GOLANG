package main

import (
	"fmt"
)

func main() {
	defer print1("Hi Harshitha")
	print2("there")
}
func print1(s string) {
	fmt.Println("Good morning", s)
}
func print2(s string) {
	fmt.Println("Did your workout today!")
}
