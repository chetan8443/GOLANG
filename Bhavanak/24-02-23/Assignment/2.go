//Write a program that declares an array of floats and calculates the average of the elements.

package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&n)
	var arr [7]float64
	fmt.Print("Enter the items in array :=")
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	var sum float64 = 0
	for i := 0; i <= n; i++ {
		sum += (arr[i])
	}
	avg := (float64)) / (float64(n))
	fmt.Println("sum =", sum)
	fmt.Println("avg =", avg)
}
