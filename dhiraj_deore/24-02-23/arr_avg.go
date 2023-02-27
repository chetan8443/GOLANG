// Write a program that declares an array of floats and calculates the average of the elements.

package main

import (
	"fmt"
)

func main() {
	arr := [5]float64{1.2, 3.3, 2.1, 4.5, 6.5} // declare array of float
	var sum float64
	count := 0

	for _, i := range arr {
		sum += i // sum of all elements
		count++ // counting elements
	}

	avg := sum / float64(count) // calculating average 

	fmt.Print(avg)
}
