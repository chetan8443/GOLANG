// Pgm to check if 2 strings are anagrams(equal no. of characters, only position is different, for eg : silent,listen)
package main

import (
	"fmt"
)

// main function
func main() {
	var str1, str2 string
	fmt.Println("Enter 2 strings")
	fmt.Scan(&str1, &str2)
	anagram(str1, str2)
}
