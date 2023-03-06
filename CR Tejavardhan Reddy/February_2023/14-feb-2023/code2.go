// write a recurssive function for the factorial of a number
package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter the number:")
	fmt.Scanf("%v", &num)
	fmt.Println("The factorial of a given number is :", fact(num))
}

func fact(num int) int {
	f := 1
	if num > 0 {
		f = num * fact(num-1)
	} else if num == 0 {
		return 1
	}
	return f
}
