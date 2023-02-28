package main

import (
    "fmt"
    "sort"
)

func main() {
    // Declare a slice of strings
    fruits := []string{"banana", "apple", "orange", "kiwi"}

    // Sort the slice in alphabetical order
    sort.Strings(fruits)

    // Print the sorted slice
    fmt.Println(fruits)
}
