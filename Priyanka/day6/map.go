package main

import "fmt"

func main() {

	m := map[string]int{
		"info":    1,
		"bellit":  2,
		"company": 3,
	}

	keys := make([]string, 0, len(m))

	for k := range m {
		keys = append(keys, k)
	}

	for i := len(keys) - 1; i >= 0; i-- {
		fmt.Printf("%s: %d\n", keys[i], m[keys[i]])
	}
}
