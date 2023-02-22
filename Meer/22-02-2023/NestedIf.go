// nested if statement
package main

import "fmt"

func main() {
	var a, b, c, large int

	fmt.Print("Enter first number: ")
	fmt.Scan(&a)
	fmt.Print("Enter second number: ")
	fmt.Scan(&b)
	fmt.Print("Enter third number: ")
	fmt.Scan(&c)

	if a > b {
		if a > c {
			large = a
		} else {
			large = c
		}
	} else {
		if b > c {
			large = b
		} else {
			large = c
		}
	}
	fmt.Println("Largest number is ", large)
}
