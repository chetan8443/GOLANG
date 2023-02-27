package module

import (
	"fmt"
	"sort"
)

// Function to give sum of elements of given array
func Sum(arr [5]int) {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Sum of elements of array=", sum)
}

// Function to give Average of elements of given array
func Avg(arr [5]float64) {
	sum := 0.0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("Average of elements of array=", sum/float64(len(arr)))
}

// Function to sort the elemets of slice of string
func Sort(str []string) {
	sort.Strings(str)
	fmt.Println(str)
}

// Function which return int slice without even numbers
func RemEven(s1 []int) {
	var s2 []int
	for _, num := range s1 {
		if num%2 != 0 {
			s2 = append(s2, num)
		}
	}
	fmt.Println("Final Slice=", s2)
}
