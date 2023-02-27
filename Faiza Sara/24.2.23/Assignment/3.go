// Write a program that declares a slice of strings and sorts it in alphabetical order.
package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"Cap", "Bag", "Mouse", "Watch", "Laptop"}
	fmt.Println("Before sorting ", words)

	sort.Strings(words)
	fmt.Println("After sorting ", words)
}
