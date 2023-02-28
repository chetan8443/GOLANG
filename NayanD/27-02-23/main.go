package main

import "fmt"

func main() {
	// Declare a map of strings to integers.
	dict := map[string]int{
		"nayan":     10,
		"dheeraj":   30,
		"arbind":    70,
		"prathmesh": 80,
	}

	// Print out the values in reverse order.
	keys := make([]string, 0, len(dict))
	for k := range dict {
		keys = append(keys, k)
	}
	for i := len(keys) - 1; i >= 0; i-- {
		k := keys[i]
		v := dict[k]
		fmt.Printf("%s: %d\n", k, v)
	}
}
