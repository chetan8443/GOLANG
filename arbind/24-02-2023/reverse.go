package main

import "fmt"

var num int

func main() { // creating main function
	fmt.Printf("Enter a number: ") // printing a statement
	fmt.Scanf("%d", &num)          // taking input from user
	var rem int                    // initialize a variable
	rev := 0                       // initialize variable to 0
	for num > 0 {                  // checking condition
		rem = num % 10     //condition
		rev = rev*10 + rem // condition
		num = num / 10     // condition
	}
	fmt.Println("The reverse number is :", rev) // printing reverse number
}
