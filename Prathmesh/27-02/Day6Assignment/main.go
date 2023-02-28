//Write a program in golang that declares a map of strings to integers and prints out the values in reverse order

package main

import (
	"fmt"
)

func main() {
	// Declare a map of strings to integers
	m := map[string]int{
		"apple":     2,
		"banana":    3,
		"cherry":    1,
		"durian":    4,
		"applellel": 5,
	}
	fmt.Println(m)
	// Create a slice to store the keys in reverse order
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	fmt.Println(keys)
	// Print out the values in reverse order
	for i := len(keys) - 1; i >= 0; i-- {
		fmt.Printf("%v: %v\n", keys[i], m[keys[i]])
	}
}
