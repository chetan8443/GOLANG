// Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import "fmt"

func main() {
	dict := map[string]int{"a": 1, "b": 2, "c": 3}
	var s []int
	for _, v := range dict {
		s = append(s, v)

	}
	for i := len(s) - 1; i >= 0; i-- {
		fmt.Println(s[i])
	}
}
