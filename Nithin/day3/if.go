// Finding the biggest number among 3 numbers
package main

import "fmt"

var a int
var b int
var c int

//main function
func main() {
	fmt.Println("enter the numbers")
	fmt.Scanf("%d%d%d", &a, &b, &c) //Taking the input from the user
	if a < b {                      //checking whether a is bigger than b
		if b > c { //checking whether b is bigger than c
			fmt.Printf("%d is the greatest number :", b)
		} else {
			fmt.Printf("%d is the greatest number :", c)
		}

	} else if a > c { //checking whether a is bigger than c
		fmt.Printf("%d is the greatest number:", a)
	} else {
		fmt.Printf("%d is the greatest number:", c)
	}
}
