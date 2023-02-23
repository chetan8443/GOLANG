//  Write a program in Golang for, How to compare two arrays equal in size or not.

package main

import "fmt"

func main() {
	fmt.Println("checking the size of two slices equal or not")

	var n int
	fmt.Println("enter the number of elements u want to add in slice1")
	fmt.Scanln(&n)
	// declaring and initializing the two slices

	var slice1 = []int{}
	var slice2 = []int{}
	// Adding elements in slices
	fmt.Println("Enter the elements: ")
	for i := 1; i <= n; i++ {
		var j int
		fmt.Scanln(&j)
		slice1 = append(slice1, j)

	}
	var p int
	fmt.Println("enter the number of elements u want to add in slice2")
	fmt.Scanln(&p)
	// Adding elements in slices
	fmt.Println("Enter the elements: ")
	for i := 1; i <= p; i++ {
		var j int
		fmt.Scanln(&j)
		slice2 = append(slice2, j)
	}
	// Calling slicequal function to check whether the slices are in equal length or not
	fmt.Println(intSlicesEqual(slice1, slice2))
}

// function for checking the slices length
func intSlicesEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	return true
}
