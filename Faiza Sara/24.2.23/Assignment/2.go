// Write a program that declares an array of floats and calculates the average of the elements.

package main

import (
	"fmt"
)

func main() {
	arr := [5]float64{2.73, 3.32, 1.14, 3.5, 2.5}
	var sum float64
	var count float64 = 0

	for _, i := range arr {
		sum += i
		count++
	}

	avg := sum / count

	fmt.Print(avg)
}
