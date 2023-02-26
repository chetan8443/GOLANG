package main

import "fmt"

func main() {
	var arr1 = [4]float32{1.2, 2.3, 4.5, 1.4}
	var i int
	var j float32 = 0.0
	fmt.Println(arr1)
	for i = 0; i < len(arr1); i++ {
		j = arr1[i] + j

	}
	fmt.Println(" average of numbers in a given array = ", j/2)
}
