// Write a program to find common elements in  two slice
package main

import "fmt"

func main() {
	s1 := []int{10, 20, 30, 40, 50, 60, 70}
	s2 := []int{30, 50, 55, 66, 880, 10, 90}
	fmt.Println("common elements :")

	for i := 0; i < len(s1); i++ {
		for k := 0; k < len(s2); k++ {
			//comaparing two slices
			if s1[i] == s2[k] {
				fmt.Println(s1[i]) //printing common elements
			}
		}
	}
}
