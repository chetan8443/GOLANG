// program that declares a slice of strings and sorts it in alphabetical order.
package main

import (
	"fmt"
	"sort"
)

func main() {
	stringSlice := []string{"Syed", "Meer", "Fahad", "Ali"}
	fmt.Println("Before Sorting :")
	for _, val := range stringSlice {
		fmt.Print(val, " ")

	}
	fmt.Println("\nAfter Sorting :")
	sort.Strings(stringSlice)
	for _, val := range stringSlice {
		fmt.Print(val, " ")

	}
}
