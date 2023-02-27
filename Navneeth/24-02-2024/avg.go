//calculates the average value of float numbers in an array

package main

import "fmt"

func main() {
	var sum float32
	var arr = [5]float32{10.4, 3.14, 5.67, 8.75, 3.98}
	var avg float32
	for i := 0; i < 5; i++ {
		sum += arr[i]
	}
	avg = sum / 5
	fmt.Printf("Average value of array: %f\n", avg)
}
