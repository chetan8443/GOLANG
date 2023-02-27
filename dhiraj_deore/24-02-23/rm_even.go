//Write a program that declares a slice of integers and removes all even numbers from the slice.
//Print out the resulting slice.

package main

import "fmt"

func main() {
	slice := []int{12, 32, 45, 21, 34, 43} // slice of integers
	fmt.Println("Slice before removing even numbers :", slice)

	var odd []int // new slice for storing result

	for _, v := range slice {
		if v%2 != 0 {
			odd = append(odd, v) // appending all odd elements into new slice and skipping even elements

		}
	}
	fmt.Println("Slice after removing even numbers :", odd)

}
