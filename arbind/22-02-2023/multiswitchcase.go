// Use case of multi switch statement

package main

import "fmt"

func main() { //main function
	fmt.Println("Welcome to multi switch case statements") //printing a statement

	day := 10    //Initialize a variable
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
