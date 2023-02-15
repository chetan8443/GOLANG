// Write a code to swap 2 numbers

package main

import "fmt"

func main() {
	fmt.Print("Enter the numbers 1 : ")
	n1 := 0
	fmt.Scanln(&n1)

	fmt.Print("Enter the numbers 2 : ")
	n2 := 0
	fmt.Scanln(&n2)

	n1, n2 = n2, n1
	fmt.Print("After swapping the numbers : ", n1, n2)
}
