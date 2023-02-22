// Doing the operations using the switch cases
package main

import "fmt"

var a int
var b int
var c int

func main() {
	fmt.Printf("enter the values")
	fmt.Scanln(&a, &b)
	fmt.Printf("press 1) for addition\n press 2) for subtract\n press 3) for multiplication")
	fmt.Scanln(&c)

	switch c {
	case 1:
		add := a + b
		fmt.Printf("added value is %d", add)
	case 2:
		subtract := a - b
		fmt.Printf("subtracted value is %d", subtract)
	case 3:
		multiply := a * b
		fmt.Printf("multiplyed value is %d", multiply)
	default:
		fmt.Printf("your given number is out of choice")
	}
}
