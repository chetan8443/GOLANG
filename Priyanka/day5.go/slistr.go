package main

import (
	"fmt"
	"sort"
)

func main() {
	fruits := []string{"apple", "banana", "orange", "kiwi", "pear"}
	sort.Strings(fruits)

	fmt.Println(fruits)
}
