//write a code to compare two numbers

package main

import "fmt"

func main() {
	var n1 int
	var n2 int
	fmt.Println("Enter two numbers to compare : ")

	fmt.Scanln(&n1)
	fmt.Scanln(&n2)

	if n1 > n2 {
		fmt.Print(n1, " is greater than ", n2)
	} else {
		fmt.Print(n2, " is greater than ", n1)
	}
}
