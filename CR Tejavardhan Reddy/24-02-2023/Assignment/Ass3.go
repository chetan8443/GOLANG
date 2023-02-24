// Write a program that declares a slice of strings and sorts it in alphabetical order.
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print("Enter the size of the array:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]string, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	sort.Strings(arr)
	fmt.Print(arr)
}
