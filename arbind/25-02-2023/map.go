package main

import (
	"fmt"
)

func main() {
	// Define the arrays to be mapped
	keys := []string{"one", "two", "three"}
	values := []int{1, 2, 3}

	// Create an empty map
	m := make(map[string]int)

	// Map the arrays to the map
	for i, key := range keys {
		m[key] = values[i]
	}

	// Print the resulting map
	fmt.Println(m)
}
