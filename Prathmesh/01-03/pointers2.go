//Creating a slice and modifying its elements using a pointer

package main

import "fmt"

func main() {
	fmt.Println("pointers")
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	sqr(&slice1) // giving address of slice to function of pointer
	fmt.Println(slice1)

}

func sqr(s *[]int) { // function takes pointer of slice of int as parameter
	for _, v := range *s {
		*s = append(*s, v*v) // appending more values to slice using pointer
	}
}
