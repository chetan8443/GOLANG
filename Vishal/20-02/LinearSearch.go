package main

import "fmt"

func main() {
	var index int
	index = linearSearch()
	if index != 404 {
		fmt.Println("Value is present at index : ", index)
	} else {
		fmt.Println("Value is Not present in array ")
	}
}

func linearSearch() int {

	arr1 := []int{12, 34, 56, 66, 76, 44, 73, 12, 38, 43}

	var input int
	fmt.Print("Enter a value to search in array: ")
	fmt.Scanln(&input)
	for i := 0; i < len(arr1)-1; i++ {
		if input == arr1[i] {
			fmt.Print(" Value found in array ")
			return i
		}
	}

	return 404
}
