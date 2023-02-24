// Write a program using for loop and using continue and break
package main

import "fmt"

func main() { // creating main function
	for i := 10; i < 20; i++ { // Iterating through for loop
		if i == 15 { // check the condition using if block
			i = i + 1
			continue //use of continue
		} else if i == 18 { //checking the condition using else if
			break // control the statement using break statement
		}
		fmt.Println("the value of i is : ", i) // Printing the value of i
	}
}
