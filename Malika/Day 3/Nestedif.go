/*
Write a program that based on NestedIf statements.
*/

// Largest of three numbers using nested if else
package main

import "fmt"

func printLargestOfThreeNumbers(a int, b int, c int) {
	if a > b {
		if a > c {
			fmt.Println("largest number is a = ", a)
		}
	} else {
		if b > c {
			fmt.Println("largest number is b = ", b)
		} else {
			fmt.Println("largest number is c = ", c)
		}
	}
}

func main() {
	printLargestOfThreeNumbers(100, 50, 20)
	printLargestOfThreeNumbers(10, 50, 20)
	printLargestOfThreeNumbers(1, 500, 20000)
}
