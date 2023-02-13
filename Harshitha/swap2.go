package main

import "fmt"

func main() {
	var a int = 90
	var b int = 6
	fmt.Println("Before Swapping")
	fmt.Println(a, b)
	var c int
	c = a
	a = b
	b = c
	fmt.Println("After Swapping")
	fmt.Println(a, b)
}
