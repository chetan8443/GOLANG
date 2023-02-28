// Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := map[string]int{"orange": 5, "apple": 7,
		"mango": 3, "strawberry": 9}

	keys := make([]string, 0, len(fruits))

	for k := range fruits {
		keys = append(keys, k)
	}
	fmt.Println(fruits)

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, k := range keys {
		fmt.Println(k, ":", fruits[k])
	}
}
