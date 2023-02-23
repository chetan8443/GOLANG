// We have to take two integers as input, and print their sum. Couple of things to note here is

// The inputs are separated by new line character
// we are solving this in golang

package main

import "fmt"

func solveMeFirst(a int, b int) int {
	return (a + b)
}

func main() {
	var a, b, res int
	fmt.Println("Enter any two numbers:")
	fmt.Scanf("%d\n%d", &a, &b)
	res = solveMeFirst(a, b)
	fmt.Println(res)
}
