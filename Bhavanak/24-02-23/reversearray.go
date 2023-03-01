// Go program to calcute the reversal array
package main

import "fmt"

func main() {
	var n, i int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&n)
	var arr [6]int
	fmt.Print("Enter the items in array :=")
	for i = 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	fmt.Println(arr)

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println("Reversal of a array:", arr)

}
