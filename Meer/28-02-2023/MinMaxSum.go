// program to calculate minimum and maximum sum of the array elements
// arr = [1,3,5,7,9]
// minSum = 1+3+5+7 = 16
// maxSum = 3+5+7+9 = 24
package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int) {
	// Write your code here
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
	arr := []int{1, 3, 5, 7, 9}
	miniMaxSum(arr)
	arr2 := []int{5, 25, 20, 15, 10}
	miniMaxSum(arr2)
}
