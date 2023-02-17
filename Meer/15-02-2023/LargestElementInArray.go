package main

import "fmt"

func main() {
	arr := [5]int{10, 20, 30, 40, 50}
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	fmt.Printf("Largest Element in the array is %v ", max)

}
