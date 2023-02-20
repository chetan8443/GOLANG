//Go program for fibonacci series

package main

import "fmt"

var c int

func main() {
	var n int
	fmt.Print("Enter the range :")
	fmt.Scanf("%v", &n)
	a := 0
	b := 1
	c := a + b
	fmt.Println("Fibonacci series")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	for i := 3; i < n; i++ {
		//c := a + b
		//fmt.Println(c)
		a = b
		b = c
		c = a + b
	}
	fmt.Print(c)
}
