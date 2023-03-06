//Creating a slice and modifying its elements using a pointer

package main

import "fmt"

func main() {
	var slice []int
	for i := 1; i <= 5; i++ {
		var a int
		fmt.Scan(&a)
		slice = append(slice, a)
	}
	fmt.Println(slice)
	modifiedSlice(&slice)
	fmt.Println(slice)

}

func modifiedSlice(slice *[]int) {
	*slice = append((*slice), 7, 8, 9)
}
