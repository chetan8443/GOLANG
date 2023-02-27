package main

import (
	"fmt"
	"sort"
)

func main() {
	dict := map[string]int{
		"prathmesh": 90,
		"dhiraj":    98,
		"nayan":     85,
		"vishal":    55,
	}
	//get the keys of the dict/map
	keys := make([]string, 0, len(dict))
	for key := range dict {
		keys = append(keys, key)

		//sort the keys
		sort.Strings(keys)
	}
	// print the sorted dick/map
	for _, key := range keys {
		fmt.Printf("%s: %d\n", key, dict[key])
	}

}
