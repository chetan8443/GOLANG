//program that declares an array of integers and prints out the sum of the elements.

package main

import "fmt"

func main() {
	var myArr = []int{10, 20, 30, 40, 50}
	sumOfArr := sumOfArrElements(myArr)
	fmt.Println("\nThe Sum of Array Elements is : ", sumOfArr)
}

func sumOfArrElements(arr []int) int {
	var sum int = 0
	fmt.Print("Array Elements are : ")
	for _, val := range arr {
		fmt.Print(val, " ")
		sum += val
	}
	return sum
}
