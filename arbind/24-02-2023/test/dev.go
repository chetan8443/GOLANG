package test

import (
	"fmt"
	"sort"
)

func Sum(arr1 [5]int) { // creating a sum function to add all the elements of an array
	fmt.Println("Wlecome to array in go lang")
	sum := 0                   // initialize a sum variable as 0
	for _, val := range arr1 { // usnig for loop iterating the array's elements
		sum = sum + val // sum all the element of array
	}
	fmt.Println("Sum of elements of array is :", sum) //printing the sum
}

func Avg(arr [5]float64) { // creating an Average function which calculated the float array average
	sum := 0.0                      //initialize
	for i := 0; i < len(arr); i++ { //using for loop iterate all the elements
		sum = +arr[i] //calculating sum

	}
	fmt.Println("Average of array is ", sum/float64(len(arr))) //printing average of array
}

func Sort(arr []string) { //creating a sort function for slice to sort the element of slice in alphabetical order
	sort.Strings(arr) // using sort method
	fmt.Println(arr)  // printing the output
}

func Even_remove(slice []int) { // creating the function to remove
	var slice1 []int            //taking a slice
	for _, val := range slice { //iterating the slice
		if val%2 != 0 { // checking the condition of even
			slice1 = append(slice1, val) // appending the elements
		}

	}
	fmt.Println(slice1) // printing the output
}
