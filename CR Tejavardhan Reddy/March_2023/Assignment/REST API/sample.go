package main

import "fmt"

func main() {
	//Reading array and printing even indexed elements
	fmt.Println("Enter the size:")
	var size int
	fmt.Scanln(&size)
	fmt.Println("Enter the elements:")
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	for i := 0; i < size; i += 2 {
		fmt.Println("The elements are", arr[i])
	}
}
