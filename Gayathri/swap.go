package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("enter a value: ")
	fmt.Scanln(&a)
	fmt.Println("enter b value: ")
	fmt.Scanln(&b)
	fmt.Println("before swapping a =",a,"b =",b)
	a=a+b
	b=a-b
	a=a-b
	fmt.Println("after swapping a =",a,"b =",b)

}