//Write a program that declares a slice of strings and sorts it in alphabetical order.

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"apple", "mango", "grapes", "banana"} // slice of strings

	fmt.Println("Slice before sorting :", slice)

	sort.Strings(slice)

	fmt.Print("slice after sorting :", slice)
}
