// program to print first duplicate element in an array
package main

import "fmt"

func main() {
	var arr [6]int
	var item int = 0
	var index int = 0
	var n int
	fmt.Println("Enter the number of elements")
	fmt.Scan(&n)

	fmt.Printf("Enter array elements: \n")
	for i := 0; i <= n; i++ {
		fmt.Scanln(&arr[i])
	}

	index = -1
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if arr[i] == arr[j] {
				item = arr[j]
				index = j
				goto OUT
			}
		}
	}

OUT:
	fmt.Println("The element which is repeating for the first time is:")
	if index != -1 {
		fmt.Println(item, index)
	} else {
		fmt.Print("There is no repeated element")
	}
}
