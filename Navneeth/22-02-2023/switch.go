// switch program to take 2 inputs and do math operation according to user input.

package main

import "fmt"

var a int //initializing variable
var b int
var opt int

//main function
func main() {
	fmt.Println("Enter two numbers to calculate.")
	fmt.Scanf("%d %d", &a, &b) //takes 2 numbers as input
	fmt.Println("Choose your operation:\n 1.addition\t 2.multiplication\t 3.division\t 4.subtraction\t")
	fmt.Scanf("%d", &opt) //chooses the operation

	//switch case
	switch opt {
	case 1: //switch case for addition
		add := a + b
		fmt.Printf("Added value is :%d", add)

	case 2: //switch case for multiplication
		mult := a * b
		fmt.Printf("Multiplicated value is: %d", mult)

	case 3: //switch case for division
		div := a / b
		fmt.Printf("Divided value is: %d", div)

	case 4: //switch case for subtraction
		sub := a - b
		fmt.Printf("Subtracted value is :%d", sub)

	default: //default case
		fmt.Printf("Invalid Case")
	}
}
