// Given an array A[] of N positive integers which can contain integers
// from 1 to P where elements can be repeated or can be absent from the array.
// Your task is to count the frequency of all elements from 1 to N.
package main

import "fmt"

func main() {
	fmt.Println("Enter the size:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Println("Enter the elements:")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	var dict = make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	//fmt.Println(dict)
	fmt.Print("Enter the P value:")
	var P int
	var res = make([]int, 0)
	fmt.Scanln(&P)
	for i := 1; i <= P; i++ {
		for _, num := range arr {
			if num == i {
				res = append(res, dict[num])
				break
			} else {
				res = append(res, 0)
				continue
			}
		}
	}
	fmt.Println(res)
}
