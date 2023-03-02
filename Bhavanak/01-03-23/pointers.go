package main

import "fmt"

func main() {
	//considering a normal variable
	var x int
	//declaring a pointer
	var p *int
	//intializing a pointer
	p = &x
	//displaying the result
	fmt.Println("Value stored in x = ", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Value stored in variable p = ", p)
	//deferencing a pointer
	fmt.Println("Value stored in x(*p) = ", *p)

	fmt.Println("Value stored in x(*p) Before Changing = ", *p)
	// changing the value of y by assigning
	*p = 500

	fmt.Println("Value stored in x(*p) after Changing = ", x)
}
