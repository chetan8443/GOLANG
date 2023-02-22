package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter any number:")
	fmt.Scanln(&n)
	if n == 0 {
		fmt.Println("You have entered '0'")
	} else if n%2 == 0 {
		fmt.Println("You have entered a even non prime number")
	} else {
		for i := 2; i < n; i++ {
			if n%i == 0 {
				fmt.Print("you have entered a odd prime number")
			} else {
				fmt.Print("you have entered a odd non prime number")
			}
		}
	}
}
