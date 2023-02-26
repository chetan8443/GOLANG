//Write a program that declares a slice of integers and removes all even numbers from the slice. Print out the resulting slice.

package main

import "fmt"

func main() {

	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9} // declaring a slice
	var new_list []int
	for _, val := range list {
		if val%2 != 0 {
			new_list = append(new_list, val) // append odd numbers to empty slice

		}
	}
	fmt.Println(new_list)
}
