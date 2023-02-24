// ALL FOUR Questions PROGRAMS CREATING WITH DIFFERNET FUNCTIONS
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("programs")
	sumofarray()
	avgofArray()
	sortArrayOfString()
	removeEvenFromSlice()
}

// 1. Write a program that declares an array of integers and prints out the sum of the elements.
func sumofarray() {
	fmt.Println("Enter the elements for an array")
	var y int = 0

	var x = [5]int{1, 2, 3, 4, 5}
	for i := 0; i <= len(x)-1; i++ {
		y += x[i]

	}
	fmt.Println("the sum of value is: ", y)
}

// Write a program that declares an array of floats and calculates the average of the elements.
func avgofArray() {

	var y int = 0

	var x = [5]int{1, 2, 3, 4, 5}
	for i := 0; i <= len(x)-1; i++ {
		y += x[i]

	}
	result := y / len(x) // avg formula
	fmt.Println("the avg of an array is: ", result)

}

// 3. Write a program that declares a slice of strings and sorts it in alphabetical order.

func sortArrayOfString() {

	var x = []string{"prathmesh", "vishal", "aniket", "dhirajD", "dhirajL"}
	sort.Strings(x) // sort function which sort the slice of an array
	fmt.Println(x)

}

//4. Write a program that declares a slice of integers and removes all even numbers from the slice. Print out the
//resulting slice.

func removeEvenFromSlice() {

	var x = []int{50, 48, 93, 17, 55, 98, 100, 60}
	var y = []int{} // declaring another slice which is empty
	var a int = len(x)
	for i := 0; i < a; i++ {
		if (i % 2) != 0 {
			y = append(y, i)
		}

	}
	fmt.Println("After removing the even number the array is : ", y)

}
