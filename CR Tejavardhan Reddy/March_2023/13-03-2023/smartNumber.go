package main

import "fmt"

func main() {
	fmt.Println("Enter a number:")
	var n int
	fmt.Scanln(&n)
	c := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			c++
		}
	}
	if n%2 == 0 {
		fmt.Println("Entered number is not a Smart Number:", n)
	} else {
		fmt.Println("Entered number is a Smart Number:", n)
	}
}
