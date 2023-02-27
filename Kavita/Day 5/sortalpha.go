package main

import (
	"fmt"
	"sort"
)

func main() {
	stringslc := []string{"e", "g", "h", "f"}

	fmt.Println("Printing the slice before sorting", stringslc)
	sort.Strings(stringslc) // inbuild function used for sorting the slice

	fmt.Println("Printing the strings after sorting", stringslc)
}
