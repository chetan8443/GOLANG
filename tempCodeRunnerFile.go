package main

import "fmt"

func main() {
	var num1, num2 int

	// Prompt user for input
	fmt.Println("Enter the first number:")
	fmt.Scanln(&num1)

	fmt.Println("Enter the second number:")
	fmt.Scanln(&num2)

	sum := num1 + num2
	fmt.Println("The sum of", num1, "and", num2, "is", sum)
}
