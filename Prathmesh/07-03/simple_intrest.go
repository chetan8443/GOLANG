package main

import "fmt"

func main() {
	var principal, time, rate int
	fmt.Print("Enter the principal amount: ")
	fmt.Scanf("%d", &principal)
	fmt.Print("Enter the time in years: ")
	fmt.Scanf("%d", &time)
	fmt.Print("Enter the rate: ")
	fmt.Scanf("%d", &rate)
	simpleInterest := (principal * time * rate) / 100 // formula to calculate simple intrest
	fmt.Println("The simple interest is: ", simpleInterest)
}
