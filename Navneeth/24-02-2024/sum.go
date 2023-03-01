//finds the sum of elements in an array

package main

import "fmt"

var sum int
var size int

func main() {
	fmt.Println("Enter size of your array.")
	fmt.Scanf("%d", &size)      //takes the size of array
	var arr = make([]int, size) //makes an array of size given by user
	for i := 0; i < size; i++ {
		fmt.Printf("Enter %d element: ", i)
		fmt.Scanf("%d", &arr[i]) //takes in elements for array
	}

	for i := 0; i < size; i++ {
		sum += arr[i] //calculates the sum
	}
	fmt.Printf("Sum of all the elements: %d\n", sum)
}
