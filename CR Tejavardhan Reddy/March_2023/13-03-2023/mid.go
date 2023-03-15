package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Enter the array size:")
	var size int
	fmt.Scanln(&size)
	var mid int
	if size%2 != 0 {
		mid = (size / 2)
	} else {
		mid = (size / 2) - 1
	}
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	if sort.IntsAreSorted(arr) {
		fmt.Println("The mid value is:", arr[mid])
	} else {
		sort.Ints(arr)
		fmt.Println("The mid value is:", arr[mid])
	}
}
