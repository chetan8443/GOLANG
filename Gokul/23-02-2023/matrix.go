// // Pgm to find sum of elements of a 2 X 2 matrix, numbers entered by the user
package main

import "fmt"

func main() {
	sum := 0
	//declaration of a 2-D matrix
	var A [2][2]int
	fmt.Println("Enter elements(2 x 2 matrix)")
	//for loop
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Scan(&A[i][j])
			sum = sum + A[i][j]
		}
	}
	fmt.Println("Sum of elements is ", sum)
}
