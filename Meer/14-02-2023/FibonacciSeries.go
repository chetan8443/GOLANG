package main

import "fmt"

func main() {
	n1 := 0
	n2 := 1
	var n3, count int

	fmt.Println("Enter the Count for fibonacci series ")
	fmt.Scanln(&count)

	fmt.Print(n1, " ", n2)

	for i := 2; i <= count; i++ {
		n3 = n1 + n2
		fmt.Print(" ", n3)
		n1 = n2
		n2 = n3
	}
}
