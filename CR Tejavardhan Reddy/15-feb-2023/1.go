//Go program for fibonacci series

package main

import "fmt"

var c int

func main() {
	var n int
	fmt.Println("Enter the range :")
	fmt.Scanf("%v", &n)
	a := 0
	b := 1
	c := a + b
	fmt.Println(a, b, c)
	for i := 3; i < n; i++ {
		//c := a + b
		//fmt.Println(c)
		a = b
		b = c
		c = a + b
	}
	fmt.Println(c)
}
