package test

import "fmt"

//var b int = 10

func Emp() {
	fmt.Println("I an an employee of infobellit")
}

func Switchcase() { //main function
	fmt.Println("Welcome to multi switch case statements") //printing a statement

	day := 4     //Initialize a variable
	switch day { // switch statement
	case 1, 3, 5: // multiple switch statement
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")

	case 6, 7:
		fmt.Println("Weekend")
	default: // default case statement
		fmt.Println("You have entered invalid day")

	}

}

func If_else() { // creating main function
	fmt.Println("Welcome to golang") // Printing a statement
	var a int = 700                  //Initialize the value of a
	if a == 200 {
		// checking the condition
		fmt.Println("Value of a  is  200") // if condition is true then
		// display the following */
	} else if a == 300 { // else if statement with given codition
		fmt.Println("Value of a is 300")
	} else if a == 400 { // else if statement with another condition
		fmt.Println("Value of a is 400")
	} else { // else statement
		fmt.Println("None of the above value matched")
	}
}
