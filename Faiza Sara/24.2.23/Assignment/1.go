// Write a program that declares an array of integers and prints out the sum of the elements.

package main

import (
	"fmt"
)

func main() {
	arr := [6]int{11, 22, 33, 44, 55, 66}
	sum := 0
	for _, i := range arr {
		sum += i
	}
	fmt.Print(sum)
}
