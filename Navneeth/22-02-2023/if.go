// Nested if program to check which is the greatest number among three inputs

package main

import "fmt"

//main function
func main() {
	var a int //initializing the variables
	var b int
	var c int
	fmt.Println("Enter three numbers.")
	fmt.Scanf("%d %d %d ", &a, &b, &c) //takes input of 3 numbers
	if a < b {                         //checks if a is greater than b
		if b > c { //checks if b is greater than c
			fmt.Printf("%d is the greatest number \n", b)
		} else { // if b is not the greatest then c is greater
			fmt.Printf("%d is the greatest number \n", c)
		}
	} else if a > c { //checks if a is the greatest
		fmt.Printf("%d is the greatest number \n", a)
	} else {
		fmt.Printf("%d is the greatest number \n", c)
	}

}
