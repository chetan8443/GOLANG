package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the number to be reversed")
	fmt.Scanln(&n)
	rem := 0
	result := 0
	for n > 0 {
		rem = n % 10
		result = result*10 + rem
		n = n / 10
	}
	fmt.Println(result)

}
