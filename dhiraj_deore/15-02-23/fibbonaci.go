// write a code for fibbonacci series

package main

import "fmt"

func main() {
	fmt.Print("Enter the number of element in the series : ")
	var n int
	fmt.Scanln(&n)

	n1 := 0
	n2 := 1
	fmt.Print(n1, " ", n2)

	for i := 0; i < n; i++ {
		n3 := n1 + n2
		fmt.Print(" ", n3)

		n1 = n2
		n2 = n3
	}
}
