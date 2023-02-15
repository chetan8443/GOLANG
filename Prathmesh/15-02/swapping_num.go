package main

// fmt package provides the function to print anything
import "fmt"

func swap_num() {
	var number1, number2 int
	number1 = 45
	number2 = 63
	fmt.Println("Numbers before swapping: \n Number 1 =", number1, "\n Number 2 =", number2)
	number1, number2 = number2, number1
	fmt.Println("Numbers after swapping:\n Number 1 =", number1, "\n Number 2 =", number2, "\n(Swap using a one-liner syntax)")
}
