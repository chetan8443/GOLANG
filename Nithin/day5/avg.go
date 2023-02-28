//calculates the average value of float numbers in an array

package main

import "fmt"

func main() {
	var sum float32
	var arr = [6]float32{10.4, 3.14, 5.67, 8.75, 3.98,10.30}
	var avg float32
	for i := 0; i < 6; i++ {
		sum += arr[i]
	}
	avg = sum / 6
	fmt.Printf("Average value of array: %f\n", avg)
}
