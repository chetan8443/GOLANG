package main

//importing custom packages
import (
	b "example/constant_var"
	a "example/fact_reccursion"
	"fmt"
)

// main function
func main() {
	// calling constant_var function
	b.Const_var()
	// calling factorial function
	result := a.Factorial(5)
	fmt.Println("factorial is:", result) //printing  factorial result
}
