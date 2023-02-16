//write a Program to check if  number is prime or not

package main

import "fmt"

func main() {
	var number int
	count := 0
	fmt.Println("Enter number : ")
	fmt.Scan(&number)
	for i := 2; i <= number/2; i++ {
		if number%i == 0 {
			count++
			break
		}
	}
	if count == 0 && number != 1 {
		fmt.Print(number, " is a prime number ")
	} else {
		fmt.Print(number, " is not prime number")
	}
}
