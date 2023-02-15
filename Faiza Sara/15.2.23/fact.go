package main

import (
	"fmt"
)

func main() {
	var a int
	f := 1
	fmt.Print("Enter a number")
	fmt.Scan(&a)
	for i := 1; i <= a; i++ {
		f = f * i
	}
	fmt.Printf("The Factorial of %d is %d", a, f)
}
