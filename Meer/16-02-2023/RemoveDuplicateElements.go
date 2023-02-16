package main

import (
	"fmt"
)

func main() {

	sliceValues := []int{15, 23, 40, 35, 23, 15, 50, 36, 40}

	fmt.Println(sliceValues)

	removeDuplicateValuesSlice := removeDuplicateValues(sliceValues)

	fmt.Println(removeDuplicateValuesSlice)
}

func removeDuplicateValues(slice []int) []int {
	keys := make(map[int]bool)
	list := []int{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
