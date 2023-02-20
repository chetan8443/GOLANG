package main

import "fmt"

func main() {
	m1 := map[string]string{
		"India":   "Delhi",
		"USA":     "WDC",
		"England": "London",
		"China":   "Shanghai",
	}

	rangeMap(m1)

	fmt.Println()

	fmt.Println("Sum of variables:", varSum(1, 2, 3, 4, 5))

	fmt.Println()

	newVar := nextInt()

	fmt.Println("Next Value:", newVar())
	fmt.Println("Next Value:", newVar())
	fmt.Println("Next Value:", newVar())
	fmt.Println("Next Value:", newVar())

	fmt.Println()

	newVar1 := nextInt()
	fmt.Println("Next Value:", newVar1())
	fmt.Println("Next Value:", newVar1())
	fmt.Println("Next Value:", newVar1())
	fmt.Println("Next Value:", newVar1())

	fmt.Println()

	num := 5
	fmt.Printf("Factorial of %d is %d", num, factorial(num))

	num1 := -2
	fmt.Printf("The sqaure root of %d is %0.2f", num, sqrt(float64(num1)))
}
