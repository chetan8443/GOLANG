//Creating a slice and modifying its elements using a pointer

package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size :")
	fmt.Scanln(&n)
	var arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Println("Enter the elements :")
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	changevalue(&arr)
}
func changevalue(arr *[]int) {
	*arr = append((*arr), 6, 7, 8, 9)
	fmt.Println(arr)
	fmt.Println(&arr)
}
