package main

import "fmt"

func main() {
	fmt.Println("Enter the Number ")
	var num int
	fmt.Scanln(&num)

	if num%2 == 0 {
		fmt.Println("Entered Number is Even")
	} else {
		fmt.Println("Entered Number is Odd")
	}
}
