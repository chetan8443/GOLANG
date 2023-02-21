// Golang program to remove duplicate
// values from Slice
package main

import (
	"fmt"
)

func removeDuplicateValues(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	// If the key(values of the slice) is not equal
	// to the already present value in new slice (list)
	// then we append it. else we jump on another element.
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
		fmt.Println(keys)
		fmt.Println(entry)
		fmt.Println(list)
	}
	return list
}

func main() {

	// Assigning values to the slice
	intSliceValues := []int{1, 2, 3, 4, 5, 2, 3, 5, 7, 9, 6, 7}

	// Printing original value of slice
	fmt.Println(intSliceValues)

	// Calling function where we
	// are removing the duplicates
	removeDuplicateValuesSlice := removeDuplicateValues(intSliceValues)

	// Printing the filtered slice
	// without duplicates
	fmt.Println(removeDuplicateValuesSlice)
}
