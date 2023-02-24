// program to calculate the positive , negative and zero numbers ratio within the array
package main

import "fmt"

func main() {

	arr := []int32{1, 1, 0, -1, -1}
	arr2 := []int32{1, 2, 3, -1, -2, -3, 0, 0}
	plusMinus(arr)
	fmt.Println("---------")
	plusMinus(arr2)

}
func plusMinus(arr []int32) {
	// Write your code here
	var positiveNo, negativeNo, zero int32
	n := int32(len(arr))
	for _, val := range arr {
		if val > 0 {
			positiveNo++
		} else if val < 0 {
			negativeNo++
		} else {
			zero++
		}

	}

	fmt.Println(float64(positiveNo) / float64(n))
	fmt.Println(float64(negativeNo) / float64(n))
	fmt.Println(float64(zero) / float64(n))

}
