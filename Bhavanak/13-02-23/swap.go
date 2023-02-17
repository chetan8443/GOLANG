package main

import "fmt"

func main() {
	var a int = 2
	var b int = 3
	fmt.Println("Before Swapping")
	fmt.Println(a)
	fmt.Println(b)
	var c int
	c = a
	a = b
	b = c
	fmt.Println("After Swapping")
	fmt.Println(a, "\n", b)
}
