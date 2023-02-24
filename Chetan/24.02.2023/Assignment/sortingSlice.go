// Write a program that declares a slice of strings and sorts it in alphabetical order.

package main

import (
	"fmt"
	"sort"
)

func main() {
	sl1 := []string{"arjun", "chetan", "rahul", "sagar", "nikhil"}	//declaring& initializing slice
//  sort.Strings(sl1)       

sort.Slice(sl1, func(i, j int) bool {
    return sl1[i] < sl1[j]
})
	fmt.Println(sl1)
}
