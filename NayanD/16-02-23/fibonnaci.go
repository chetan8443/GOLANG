// write a program for fibbonacci series

package main

import "fmt"

func main() {

	fmt.Print("Enter the number  : ")
	var a int
	fmt.Scanln(&a)

	a1 := 0
	a2 := 1
	fmt.Print(a1, " ", a2)

	for i := 0; i < a; i++ {
		a3 := a1 + a2
		fmt.Print(" ", a3)

		a1 = a2
		a2 = a3
	}
}
