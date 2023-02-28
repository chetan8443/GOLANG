//Write a program that declares a map of strings to integers and prints out the values in reverse order

package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{
		"Ritesh":  33,
		"Kumar":   31,
		"MOhanty": 67,
	}
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	for _, k := range keys {
		fmt.Println(k, m[k])
	}

}
