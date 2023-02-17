// factorial of a number using recursion
package main

import "fmt"

func fact(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	fmt.Println(fact(n))
}
