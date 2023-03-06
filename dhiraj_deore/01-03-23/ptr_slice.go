//Creating a slice and modifying its elements using a pointer

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	sqr(&slice) // giving address of slice to function as pointer 
	fmt.Println(slice)

}

func sqr(s *[]int) { // function which takes pointer of slice of int as parameter
	for _, v := range *s {
		*s = append(*s, v*v) // appending more values to slice using pointer
	}
}
