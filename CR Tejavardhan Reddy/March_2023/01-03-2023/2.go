//Golang code to read an array of integers and finds the difference between max and min

package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print("Enter the size:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	sort.Ints(arr)
	minimum := arr[0]
	maximum := arr[size-1]
	fmt.Print("The difference between minimum and maximum is:", maximum-minimum, arr)
}

// Input:
//Enter the size:5
//
//Output:
// 5
// 43
// 2
// 35
// 67
