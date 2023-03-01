package main

import "fmt"

func findArraySum(arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	return sum
}

func main() {
	fmt.Println(findArraySum([]int{1, 2, 3, 4, 5}))
	fmt.Println(findArraySum([]int{3, 5, 7, 2, 1}))
	fmt.Println(findArraySum([]int{9, 8, 6, 1, 0}))
}
