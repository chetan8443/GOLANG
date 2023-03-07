// program to count the number of words in a variable i.e. in camel case
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "oneTwoThree"
	s1 := "saveChangesInTheEditor"
	r := camelcase(s)
	fmt.Printf("Number of Words in %s are %d \n", s, r)
	r = camelcase(s1)
	fmt.Printf("Number of Words in %s are %d \n", s1, r)
}

func camelcase(s string) int {
	// Write your code here
	count := 1
	str := strings.ToUpper(s)
	for i := 0; i < len(s); i++ {
		if s[i] == str[i] {
			count++
		}
	}
	return count
}
