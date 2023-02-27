// Write a program that declares an array of integers and prints out the sum of the elements.

package main

import (
	"fmt"
)

func main() {
	arr := [6]int{12, 32, 45, 21, 34, 43} // declare array
	sum := 0
	for _, i := range arr {
		sum += i // sum of all elements of array
	}
	fmt.Print(sum)
}
