package main

import "fmt"

func main() {
	var num int = 10

	if num > 0 {
		fmt.Println("The number is positive")

		if num%2 == 0 {
			fmt.Println("The number is even")
		} else {
			fmt.Println("The number is odd")
		}
	} else if num == 0 {
		fmt.Println("The number is zero")
	} else {
		fmt.Println("The number is negative")
	}
}
