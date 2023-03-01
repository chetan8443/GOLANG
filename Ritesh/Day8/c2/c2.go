//Creating a slice and modifying its elements using a pointer

package main

import "fmt"

type Item struct {
	Value int
}

func alter(t *Item) {
	(*t).Value = 100
}

func main() {
	items := []Item{{0}, {1}, {2}}

	for i := range items {
		alter(&items[i])
	}

	fmt.Println(items)
}
