//Write a program that declares an array of floats and calculates the average of the elements.

package main

import "fmt"

func main() {

	var sum float64 = 0
	var average float64 = 0

	arr := [...]float64{1.2, 3.5, 2.11, 4.5, 1.0} // declaring an integer array

	for i := 0; i <= len(arr)-1; i++ {
		sum = sum + arr[i]

		average = sum / float64(len(arr))
	}

	fmt.Println("sum of array elements: ", average) //printing average of elements in array

}
