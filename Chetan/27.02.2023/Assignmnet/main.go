//Write a program in golang that declares a map of strings to integers and prints out the values in reverse order

package main

import (
	"fmt"
)

func main() {
	// Declare a map of strings to integers
	m := map[string]int{
		"red":    3,
		"yellow": 1,
		"black":  4,
		"white":  2,
		"orange": 5,
	}
	// printing before reverse order
	fmt.Println(m)
	// Create a slice to store the keys in reverse order
	values := make([]int, 0, len(m))


	for _, v := range m {
		values = append(values, v)
	}




	// Print out the values in reverse order
	fmt.Println("VALUES IN REVERSE ORDER : ")
	for i := len(values) - 1; i >= 0; i-- {
		
		fmt.Print(values[i]," ")
	}
}
