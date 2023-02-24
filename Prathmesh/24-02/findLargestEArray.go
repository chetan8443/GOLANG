// Golang program to find the largest elements
// from the array

package main

import "fmt"

func main() {
	var large int = 0
	var arr = [5]int{20, 85, 955, 897, 2}

	large = arr[0] //large variable for comparing with other elements
	for i := 1; i <= 4; i++ {
		if large < arr[i] {
			large = arr[i]
		}
	}

	fmt.Println("Largest element is: ", large)
}
