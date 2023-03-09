// program that declares an array of floats and calculates the average of the elements.
package main

import "fmt"

func main() {
	var floatArr = []float32{13.53, 17.743, 23.9087}
	avg := avgOfFloatElements(floatArr)
	fmt.Println("\nThe Average of Array Elements is : ", avg)
}

func avgOfFloatElements(arr []float32) float32 {
	n := len(arr)
	var total float32
	fmt.Print("Array Elements are : ")
	for _, val := range arr {
		fmt.Print(val, " ")
		total += (val)
	}
	return (total / float32(n))
}
