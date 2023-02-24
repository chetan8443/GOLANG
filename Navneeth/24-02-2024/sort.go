//sorts string in slice in alphabetical order

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"a", "c", "q", "b", "l"}

	sort.Strings(slice)
	fmt.Println(slice)
}
