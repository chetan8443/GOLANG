package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("programs")
	sum() // 1. Write a program that declares an array of integers and prints out the sum of the elements.
	avg()
	fmt.Println()
	sortString()
	fmt.Println()
	passOdd()
}

func sum() {

	var sumOfEle int = 0

	var arr = [6]int{10, 20, 30, 40, 50, 60}
	for i := 0; i <= len(arr)-1; i++ {
		sumOfEle += arr[i]

	}
	fmt.Printf("The of all element is %d: ", sumOfEle)
}

// Write a program that declares an array of floats and calculates the average of the elements.
func avg() {

	var temp int = 0

	var arr = [6]int{10, 20, 30, 40, 50, 60}
	for i := 0; i <= len(arr)-1; i++ {
		temp += arr[i]

	}
	result := temp / len(arr)
	fmt.Printf("The average is %d: ", result)

}

// 3. Write a program that declares a slice of strings and sorts it in alphabetical order.
func sortString() {

	var str = []string{"xyz", "pqr", "abc", "cdf", "rti", "one"}
	sort.Strings(str)
	fmt.Println(str)

}

//4. Write a program that declares a slice of integers and removes all even numbers from the slice. Print out the

func passOdd() {

	var arr = []int{1, 4, 5, 2, 7, 9, 11, 20, 16}
	var oddArray = []int{} // declaring another slice which is empty
	var a int = len(arr)
	for i := 0; i < a; i++ {
		if (i % 2) == 0 {
			oddArray = append(oddArray, arr[i])
		}
		

	}
	fmt.Printf("After removing even numbers the array is %d: ", oddArray)

}
