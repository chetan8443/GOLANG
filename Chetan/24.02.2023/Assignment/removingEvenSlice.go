// Write a program that declares a slice of integers and removes all even numbers from the slice. Print out the
// resulting slice.

package main

import "fmt"

func main() {
	sl2 := []int{1, 2, 3, 4, 5, 6, 7}
	var sl3 []int
	for i := 0; i < len(sl2); i++ {
		if sl2[i]%2 != 0 {
			
			sl3=append(sl3, sl2[i])
		}
	}
	fmt.Println(sl3)

}
