//program to calculate minimum and maximum sum of the array elements

package main

import (
	"fmt"
	"sort"
)

func minMaxSum(arr []int) {
	sort.Ints(arr)
	var sum int
	for i := range arr {
		sum += arr[i]
	}
	min := sum - arr[len(arr)-1]
	max := sum - arr[0]
	fmt.Println(min, max)
}

func main() {
	arr := []int{1, 2, 3, 6, 9}
	minMaxSum(arr)
	arr2 := []int{2, 4, 6, 8, 10}
	minMaxSum(arr2)
}
