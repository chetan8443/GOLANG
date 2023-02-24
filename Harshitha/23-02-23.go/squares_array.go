package main

import (
	"fmt"
)

func main() {
	var size int
	fmt.Println("Enter the size of the array")
	fmt.Scanln(&size)
	var array = make([]int, size)
	var i int
	fmt.Println("Enter the elements into the array")
	for i = 0; i < size; i++ {
		fmt.Println("enter array of", i)
		fmt.Scanln(&array[i])
	}
	fmt.Println(array)
	for i = 0; i < size; i++ {
		array[i] = array[i] * array[i]
	}
	fmt.Println(array)
}
