//Write a program that declares an array of integers and prints out the sum of the elements.

package main

import "fmt"

func main() {

	var sum int = 0
	arr := [...]int{10, 30, 20, 30, 10} // declaring an integer array

	for i := 0; i <= len(arr)-1; i++ {
		sum = sum + arr[i]
	}

	fmt.Println("sum of array elements: ", sum) //printing sum of elements in array

}
