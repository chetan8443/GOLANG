package main

import "fmt"

func main() {
	fmt.Println("Enter two numbers to swap")
	var num1, num2 int
	fmt.Scan(&num1, &num2)
	fmt.Println("-------Before Swapping-------")
	fmt.Println(num1, " ", num2)

	num1 = num1 + num2
	num2 = num1 - num2
	num1 = num1 - num2

	fmt.Println("-------After Swapping-------")
	fmt.Println(num1, " ", num2)
}
