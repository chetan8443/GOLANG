//2)Write a program for multiple switch statement.

package main

import "fmt"

func main() {
	var marks int
	fmt.Println("Enter your marks :")
	fmt.Scanf("%d", &marks)

	switch marks {

	case 40, 50:
		fmt.Println("your grade is C")
	case 60, 70:
		fmt.Println("your grade is B")
	case 80, 90:
		fmt.Println("your grade is A")
	default:
		fmt.Println("Please enter valid marks")
	}
}
