// Program to rotate the array elements based on the user how many times he wants to rotate
package main

import "fmt"

//var arr = []int{1, 2, 3, 4, 5}

func main() {
	var n int
	fmt.Println("Enter the number of elements you want to enter")
	fmt.Scan(&n)
	var arr = make([]int, n)

	fmt.Println("Enter the numbers :")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println("Array Before Rotataion :")
	fmt.Println(arr)
	fmt.Println("Enter the number of roation you want in the array :")
	var r int
	fmt.Scan(&r)
	arr = rotate(arr, r)
	fmt.Println("Array After Rotataion :")
	fmt.Println(arr)

}

func rotate(arr []int, r int) []int {
	var x = arr[len(arr)-1]
	for i := len(arr) - 1; i > 0; i-- {
		arr[i] = arr[i-1]
	}
	arr[0] = x
	r--
	if r != 0 {
		rotate(arr, r)
	}
	return arr
}
