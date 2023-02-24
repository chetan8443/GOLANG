/*
Go Program to Put Positive and Negatives in a Separate Array

input :
Enter the Positive Negative Array Size = 5
Enter the Positive Negative Array Items  = 10 -99 -87 22 0
output :
The Total Number of Positive Numbers =  3
The Positive Array Elements are = 10 22 0 The Total Number of Negative Numbers =  2
The Negative Array Elements are = -99 -87
*/
package main

import "fmt"

func printArray(posNegArr []int, size int) {
	for i := 0; i < size; i++ {
		fmt.Print(posNegArr[i], " ")
	}
}
func main() {
	var size, i int

	fmt.Print("Enter the Positive Negative Array Size = ")
	fmt.Scan(&size)

	posNegArr := make([]int, size)
	posArr := make([]int, size)
	NegArr := make([]int, size)

	fmt.Print("Enter the Positive Negative Array Items  = ")
	for i = 0; i < size; i++ {
		fmt.Scan(&posNegArr[i])
	}

	positiveCount := 0
	negativeCount := 0

	for i = 0; i < size; i++ {
		if posNegArr[i] >= 0 {
			posArr[positiveCount] = posNegArr[i]
			positiveCount++
		} else {
			NegArr[negativeCount] = posNegArr[i]
			negativeCount++
		}
	}
	fmt.Println("The Total Number of Positive Numbers = ", positiveCount)
	fmt.Print("The Positive Array Elements are = ")
	printArray(posArr, positiveCount)
	fmt.Println("The Total Number of Negative Numbers = ", negativeCount)
	fmt.Print("The Negative Array Elements are = ")
	printArray(NegArr, negativeCount)
}
