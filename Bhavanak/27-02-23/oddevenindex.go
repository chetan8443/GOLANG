// Go Program to reverse an array

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of an array :=")
	fmt.Scanln(&n)
	var arr [7]int
	fmt.Println("Enter the elements of an array :")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)

	fmt.Println("The List of Array Items in Odd Index Position = ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}

	fmt.Println("The List of Array Items in Even Index Position = ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Println(arr[i])
	}

}
