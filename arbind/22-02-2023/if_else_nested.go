// Program on if-else ladder
package main

import "fmt"

func main() { // creating main function
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
