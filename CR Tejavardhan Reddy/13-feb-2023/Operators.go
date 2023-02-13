package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	arr1 := []int{}
	arr2 := []int{}
	for _, num := range arr {
		if num%2 == 0 {
			arr1 = append(arr1, num)
		} else {
			arr2 = append(arr2, num)
		}
	}
	fmt.Print("Even numbers are:", arr1, "\nOdd numbers are:", arr2)
}
