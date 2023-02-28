// Read an array from input of any size and print the difference in the
// even sum and odd sum with respect to the totalsum
package main

import (
	"fmt"
	ref "teja/sum"
)

func main() {
	fmt.Print("Enter the size:")
	var size int
	fmt.Scanf("%d", &size)
	var arr = make([]int, size)
	fmt.Println("Enter the elemenets:")
	for i := 0; i < 5; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("The entered array is:", arr)
	totalSum := ref.Sum(arr)
	evenSum := ref.EvenSum(arr)
	oddSum := ref.OddSum(arr)
	fmt.Println("The sum of elements is:", totalSum)
	fmt.Println("The sum of even elements with respect to total sum:", totalSum-evenSum)
	fmt.Println("The sum of odd elements with respect to total sum:", totalSum-oddSum)
}
