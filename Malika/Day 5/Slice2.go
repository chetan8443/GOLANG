package main

import "fmt"

func printOnlyOddNumberSlice() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 {
			slice = append(slice, arr[i])
		}
	}
	fmt.Println("All numbers:", arr)
	fmt.Println("Only Odd numbers:", slice)
}

func main() {

	printOnlyOddNumberSlice()
}
