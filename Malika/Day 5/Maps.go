// Reverse Map print
package main

import "fmt"

type Pair struct {
	Key   string
	Value int
}

func MapOfStringAndInt() {
	var nameMap map[string]int
	nameMap = map[string]int{"Adam": 1, "Bob": 2, "Carol": 3}

	var pairs []Pair

	fmt.Println("Original Map order:")
	for key, value := range nameMap {
		fmt.Println(key, "=>", value)
		pairs = append(pairs, Pair{key, value})
	}

	fmt.Println("\nReverse Map order:")
	for i := len(pairs) - 1; i >= 0; i-- {
		fmt.Println(pairs[i].Key, "=>", pairs[i].Value)
	}
}

func main() {
	MapOfStringAndInt()
}
