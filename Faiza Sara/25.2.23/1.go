// Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import (
	"fmt"
	"sort"
)

func main() {
	veggies := map[string]int{"Potato": 8, "Tomato": 7,
		"onion": 3, "Brinjal": 5}

	keys := make([]string, 0, len(veggies))

	for k := range veggies {
		keys = append(keys, k)
	}
	fmt.Println(veggies)

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, k := range keys {
		fmt.Println(k, veggies[k])
	}
}
