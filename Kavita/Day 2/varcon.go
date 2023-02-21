package main

import "fmt"

var d float32 = 0.0001
var e int = 10

func main() {

	var a string = "first"
	b := "second"
	c := 10
	const name string = "john"
	const age = 50
	const retired = false

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d, e)
	fmt.Printf("%v: %T\n", name, name)
	fmt.Printf("%v: %T\n", age, age)
	fmt.Printf("%v: %T\n", retired, retired)
}
