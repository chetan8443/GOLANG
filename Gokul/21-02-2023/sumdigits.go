//Program to print the sum of digits of a number entered by the user.
package main

import "fmt"

//Main function
func main() {
	var num int
	fmt.Println("Enter the number")
	fmt.Scan(&num)
	fmt.Print("The sum of digits of ",num," is ",digsum(num))
}

//Function to calculate the sum of digits, and return the same
func digsum(n int) int {
	sum := 0
	for n >0 {
		sum = sum + n%10
		n   = n/10
	}
	return sum
}