package main

import "fmt"

func main() {
	// Declare an array of 5 floats
	nums := [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}

	// Calculate the sum of the elements
	var sum float64
	for _, num := range nums {
		sum += num
	}

	// Calculate the average of the elements
	avg := sum / float64(len(nums))

	// Print the result
	fmt.Printf("The average of the array is: %.2f\n", avg)
}
