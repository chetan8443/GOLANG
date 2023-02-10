package main

import "fmt"

func main() {
	var num int = 0

	fmt.Printf("Enter number: ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Printf("Number is EVEN")
	} else {
		fmt.Printf("Number is ODD")
	}
}
