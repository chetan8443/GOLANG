//Write a program that declares a slice of strings and sorts it in alphabetical order.

package main

import (
	"fmt"
	"sort" // importing sort function
)

func main() {
	list := []string{"b", "e", "d", "f", "a", "c"} // declaring slice of Strings
	fmt.Println("Before sorting", list)
	sort.Strings(list) // sorting the slice using sort.Strings
	fmt.Println("After sorting", list)
}
