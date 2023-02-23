package main

import "fmt"

func main(){
	var n,a int
	fmt.Println("enter n")
	fmt.Scanln(&n)
	a=Factorial(n)
	fmt.Println(a)
}

//function to calculate factorial using reccursion
func Factorial(n int) int {

	if n == 0 {
		return 1
	}
	return n * Factorial(n-1) // factorial function calling itself
}