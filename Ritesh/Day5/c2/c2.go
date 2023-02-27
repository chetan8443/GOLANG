// 2.Write a program that declares an array of floats and calculates the average of the elements.
package main

import "fmt"

func main() {
	var a [5]float64
	a[0] = 1.3
	a[1] = 43.5
	a[2] = 21.6
	a[3] = 8.9
	a[4] = 90.7
	fmt.Println(a)

	var avg float64
	var t float64
	for _, v := range a {
		t = t + v
		avg = t / float64(len(a))
	}
	fmt.Printf("total of array is %f\n", t)
	fmt.Printf("average of array is %f\n", avg)
}
