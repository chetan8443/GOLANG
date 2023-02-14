package main

import "fmt"

var c int

func main() {
	var a int = 0
	var b int = 1
	var x int = 10
	for i := 0; i < x; i++ {
		c = a + b
		fmt.Println(a)
		a = b
		b = c
	}
}
