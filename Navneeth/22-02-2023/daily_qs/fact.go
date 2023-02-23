//this program prints factorial of a number that is given by the user

package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number.")
	fmt.Scanf("%d", &a)
	fmt.Println(fact(a))
}

func fact(a int) int {
	if a == 1 || a == 0 {
		return a
	}
	return a * fact(a-1)
}
