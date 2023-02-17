package main

import "fmt"

func main() {
	arr := []int{78, 34, 1, 3, 90, 34, -1, -4, 6, 55, 20, -65}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			tmp := 0
			if arr[i] > arr[j] {
				tmp = arr[i]
				arr[i] = arr[j]
				arr[j] = tmp
			}
		}

	}
	//fmt.Print(arr)

	for i := 0; i < len(arr); i++ {
		fmt.Print(" ", arr[i])
	}
}
