package main

import "fmt"

func main() {
	// Declare an array of integers
	nums := [5]int{1, 2, 3, 4, 5}

	// Calculate the sum of the elements
	sum := 0
	for _, num := range nums {
		sum += num
	}

	// Print the sum
	fmt.Println("Sum of the elements:", sum)
}
