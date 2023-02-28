package main

import "fmt"

func main() {
	// Declare a slice of integers.
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// Remove all even numbers from the slice.
	odds := make([]int, 0)
	for _, num := range nums {
		if num%2 != 0 {
			odds = append(odds, num)
		}
	}

	// Print out the resulting slice.
	fmt.Println(odds)
}
