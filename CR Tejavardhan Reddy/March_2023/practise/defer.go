package main

import "fmt"

func main() {
	x := 10
	if x >= 10 {
		defer print1("Hi")
	}
	defer print2("Hello")
	defer print2("MG")
	defer fmt.Println("LAst")
}
func print1(s string) {
	fmt.Println(s)
}
func print2(s string) {
	fmt.Println(s)
}
