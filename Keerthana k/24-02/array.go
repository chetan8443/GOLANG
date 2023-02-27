package main

import "fmt"

func main() {
	var arr1 = [4]int{5, 10, 5, 5}
	var i int
	var j int = 0
	fmt.Println(arr1)
	j = 0
	for i = 0; i < len(arr1); i++ {
		j = arr1[i] + j

	}
	fmt.Println("sum of numbers in a given array = ", j)
}
