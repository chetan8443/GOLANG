package main

import "fmt"

func main() {
	const myConst int = 28
	fmt.Println("The number of days in the month of February 2023 is", myConst)
	// print the value of int

	const myConst1 float64 = 1.5
	fmt.Println("The number of leaves can be taken in one month:", myConst1)
	// print the value of float64

	// Declare and initialize a constant of type bool
	const myConstant2 bool = true
	// Print the value of the constant
	fmt.Println("The value of myConstant of data type bool is:", myConstant2)

	const myConstant3 string = "Welcome to the Golang Training"
	// declare and initialize a constant string

	fmt.Println(myConstant3)
	// print the value of the constant
}
