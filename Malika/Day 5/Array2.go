//Write a program that declares an array of integers and prints out the sum of the elements.//

package main

import "fmt"

func array(arr []float32) {

	var sum float32 = 0
	var average float32 = 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	average = sum / float32(len(arr))
	fmt.Println(average)
}

func main() {
	arr := []float32{10.5, 20.7, 30.45, 40.54, 50.34}
	array(arr)
}
