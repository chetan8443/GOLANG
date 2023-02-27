// Write a program that declares an array of floats and calculates the average of the elements.
package main

import "fmt"

func main() {
	fmt.Print("Enter the size of the array:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]float64, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	var sum float64
	for _, element := range arr {
		sum += element
	}
	var average float64
	average = float64(sum) / float64(size)
	fmt.Printf("%.2f", average)
}
