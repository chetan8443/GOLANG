/*
Go program to read an array and counts the positive and negative
numbers and prints the maximum count
Input:[1,2,3,0,0,-1,-2,-3,-4]
Output:4
*/

package main

import "fmt"

func main() {
	//var nums int
	var size int
	fmt.Print("Enter the size of the array")
	fmt.Scanln(&size)
	var nums = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Print("Enter the elements")
		fmt.Scanln(&nums[i])
	}
	var negaviteCount int = 0
	var positiveCount int = 0
	for _, num := range nums {
		if num < 0 {
			negaviteCount++
		} else if num > 0 {
			positiveCount++
		}
	}
	if negaviteCount > positiveCount {
		fmt.Print(negaviteCount)
	} else {
		fmt.Print(positiveCount)
	}
}
