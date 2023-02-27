// program that declares a slice of integers and removes all even numbers from the slice.
// Print out the resulting slice.
package main

import "fmt"

func main() {
	intSlice := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println("Before Removing Even No : ")
	for _, val := range intSlice {
		fmt.Print(val, " ")
	}
	removedEvenNoSlice := removeEvenNo(intSlice)
	fmt.Println("\nAfter Removing Even No : ")
	for _, val := range removedEvenNoSlice {
		fmt.Print(val, " ")
	}
}

func removeEvenNo(intSlice []int) []int {
	var integerSlice []int
	for _, val := range intSlice {
		if val%2 != 0 {
			integerSlice = append(integerSlice, val)
		}
	}
	return integerSlice
}
