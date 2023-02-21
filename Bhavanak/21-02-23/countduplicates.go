// Count the duplicates numbers in array using Function

package main

import "fmt"

var dupcount, i int

func countDuplicates(dupArr []int, dupsize int) int {
	dupcount = 0
	for i = 0; i < dupsize; i++ {
		for j := i + 1; j < dupsize; j++ {
			if dupArr[i] == dupArr[j] {
				dupcount++
				break
			}
		}
	}
	return dupcount
}
func main() {
	var dupsize int

	fmt.Print("Enter the Even Array Size = ")
	fmt.Scan(&dupsize)
	dupArr := make([]int, dupsize)
	fmt.Print("Enter the Even Array Items  = ")
	for i = 0; i < dupsize; i++ {
		fmt.Scan(&dupArr[i])
	}
	dupcount = countDuplicates(dupArr, dupsize)
	fmt.Println("\nThe Total Number of Duplicates in dupArr = ", dupcount)
}
