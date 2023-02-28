//sorts string in slice in alphabetical order

package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []string{"a", "i","e","u","o"}

	sort.Strings(slice)
	fmt.Println(slice)
}