package main

import "fmt"

                                             // Function to check if one string is a rotation of another
func isRotation(str1, str2 string) bool {
	// If strings have different lengths, return false
	if len(str1) != len(str2) {
		return false
	}

	
	concat := str1 + str1           // Concatenate str1 with itself to check for rotations

	// Check if str2 is a substring of the concatenated string
	if (len(concat) >= len(str2)) && (concat[:len(str2)] == str2) {
		return true
	}

	return false
}

func main() {
	// Example input strings
	str1 := "hello world"
	str2 := "worldhello "

	
	if isRotation(str1, str2) {        // Check if str2 is a rotation of str1
		fmt.Printf("%s is a rotation of %s\n", str2, str1)
	} else {
		fmt.Printf("%s is not a rotation of %s\n", str2, str1)
	}
}
