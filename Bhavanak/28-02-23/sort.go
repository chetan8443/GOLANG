// Golang Program to sort usimg Bubblesort algorithm
package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the size of array :=")
	fmt.Scanln(&n)
	var arr = make([]int, n)
	fmt.Println("Enter the elements of an array :=")
	for i := 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	fmt.Println(Bubblesort(arr))
}

func Bubblesort(arr []int) []int {
	var temp int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if arr[i] < arr[j] {
				temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
