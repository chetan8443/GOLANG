// Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import "fmt"

func main() {
	dict := map[string]int{"a": 1, "b": 2, "c": 3}
	var s []int
	for _, v := range dict {
		s = append(s, v)

	}
	//fmt.Print(s)
	fmt.Print(reverse(s))
}
func reverse(s []int) []int {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
	}
	return s
}
