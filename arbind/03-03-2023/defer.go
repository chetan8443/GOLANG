package main

import (
	"fmt"
)

func main() {
	defer print1("Hi...")
	print2("there")
}
func print1(s string) {
	fmt.Println(s)
}
func print2(s string) {
	fmt.Println(s)
}
