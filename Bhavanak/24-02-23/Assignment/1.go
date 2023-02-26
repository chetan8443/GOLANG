//Write a program that declares an array of integers and prints out the sum of the elements.

package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&n)
	var arr [6]int

	fmt.Println("enter a number")
	for i = 0; i < n; i++ {
		fmt.Scanln(&arr[i])
	}
	var sum int = 0
	for i = 0; i < n; i++ {
		sum = sum + arr[i]

	}
	fmt.Println(sum)
}
