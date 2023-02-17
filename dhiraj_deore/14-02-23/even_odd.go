//write a program to check given number is even or odd

package main

import "fmt"

func main() {

	var number int
	fmt.Println("Please enter the number : ")

	fmt.Scanln(&number)

	// logic to check that number is divisible by 2 or not
	if number%2 == 0 {
		fmt.Println(number, "The given number is even number")
	} else {
		fmt.Println(number, "the given number is odd number")
	}
}
