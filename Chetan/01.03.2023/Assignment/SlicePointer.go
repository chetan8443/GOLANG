//Creating a slice and modifying its elements using a pointer

package main

import (
	"fmt"
	
)

func main() {
	slice := []string{"red","orange","black"}
	fmt.Println(slice)
	modify(&slice) // giving address of slice to function as pointer
	fmt.Println(slice)

}

func modify(s *[]string) { // function which takes pointer of slice of string as parameter

	for _, v := range *s {
		
		*s = append(*s, v) // appending more values to slice using pointer
	}
}
