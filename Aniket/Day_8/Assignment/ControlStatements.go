package main

import "fmt"

func main() {

	// Code block for "Nested if-else"
	fmt.Println("Enter your age:")
	var age int
	fmt.Scanln(&age)
	ageSwitch(age)

	// Code block for "Multiple case Switch Statements"
	fmt.Println("Enter Number between 1 to 10")
	var num int
	fmt.Scanln(&num)
	oddEven(num)

	// Code block for "Nested for loop"
	var arr = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	array(arr)
}
