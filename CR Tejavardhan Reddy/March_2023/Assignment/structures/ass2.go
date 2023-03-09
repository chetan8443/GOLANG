// Creating a slice and modifying its elements using a pointer
package main

import "fmt"

func main() {
	fmt.Print("Enter the size:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	modifyArray(&arr)

}

func modifyArray(arr *[]int) {
	(*arr)[0] = 10
	fmt.Print(arr)
}
