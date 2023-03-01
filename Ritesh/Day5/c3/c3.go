// 3. Write a program that declares a slice of strings and sorts it in alphabetical order.
package main

import (
	"fmt"
	"sort"
)

func main() {
	var a = []string{"d", "a", "v", "c", "d"}
	fmt.Printf("elements before sorting %s\n", a)
	sort.Strings(a)
	fmt.Printf("elements after sorting %s\n", a)

}
