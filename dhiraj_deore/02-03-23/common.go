// Write a program to find common elements in  two slice
package main

import "fmt"

func main() {
	slice1 := []int{12, 32, 34, 56, 76, 43, 77}
	slice2 := []int{32, 54, 45, 67, 87, 12, 77}
	fmt.Println("common elements are :")

	for i := 0; i < len(slice1); i++ {
		for k := 0; k < len(slice2); k++ {
			if slice1[i] == slice2[k] {
				fmt.Println(slice1[i])
			}
		}
	}
}
