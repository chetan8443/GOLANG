package main

import "fmt"
import a "hello/Fc"

  

func main() {
	
	str1 := "hello world"
	str2 := "worldhello "

	
	if a.IsRotation(str1, str2) {        // Check if str2 is a rotation of str1
		fmt.Printf("%s is a rotation of %s\n", str2, str1)
	} else {
		fmt.Printf("%s is not a rotation of %s\n", str2, str1)
	}
}
