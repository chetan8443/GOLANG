package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, i int
	fmt.Println("enter the size of the array")
    fmt.Scanln(&n)
    var animals []string
    for i=0;i<n;i++{
		fmt.Println("Enter ith element")
		fmt.Scanln(&animals[i])
	}
    fmt.Println("Before sorting =", animals)
    fmt.Println("length of animals=", len(animals))
    fmt.Println("length of animals=", cap(animals))
    sort.Strings(animals)
    fmt.Println("After sorting =", animals)
}

