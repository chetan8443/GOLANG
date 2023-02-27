//1.Write a program that declares an array of integers and prints out the sum of the elements.

package main

import "fmt"

func main() {
	var a [5]int

	a[0] = 1
	a[1] = 43
	a[2] = 21
	a[3] = 8
	a[4] = 90
	fmt.Println(a)
	t := 0
	for _, v := range a {
		t = t + v
	}
	fmt.Println(t)

}
