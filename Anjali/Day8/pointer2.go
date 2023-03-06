package main

import "fmt"

func main() {
	// create a slice of integers
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", nums)

	// get a pointer to the first element of the slice
	ptr := &nums[0]

	// modify the third element of the slice using the pointer
	*ptr = 10
	fmt.Println("Modified slice:", nums)
}
