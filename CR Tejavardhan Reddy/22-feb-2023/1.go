/*Write a Golang code to sort the given array
Note: The array should be user read format
Input:[88,6,3,21,4,1]
Output:[1,3,4,6,21,88]
*/

package main

import "fmt"

func main() {
	fmt.Printf("Enter the size of the array:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Println("Enter the elements into the array")
		fmt.Scanln(&arr[i])
	}
	fmt.Print(arr)
	for j := 0; j < size-1; j++ {
		for k := 0; k < size-1; k++ {
			if arr[k+1] < arr[k] {
				temp := arr[k]
				arr[k] = arr[k+1]
				arr[k+1] = temp
			}
		}
	}
	fmt.Println(arr)
}
