package main

import "fmt"

func main() {
	var a int
	fmt.Println("enter a my dear harshitha")
	fmt.Scanln(&a)
	var b int
	fmt.Println("enter b darling")
	fmt.Scanln(&b)
	fmt.Println("The Arithmetic opeartions are:")
	fmt.Println("The Addition of two numbers", a, "and", b, "is")
	fmt.Println(a + b)
	fmt.Println("The Increment is")
	a++
	fmt.Println(a)
	fmt.Println("The Addition of two numbers", a, "and", b, "is")
	fmt.Println(a + b)
}
