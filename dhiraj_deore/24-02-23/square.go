//Write a program to calculate square the elements in slice:

package main

import "fmt"

func main() {

	slice := []int{2, 3, 4, 5, 6}

	for _, v := range slice {
		fmt.Println("Element :", v, " Square :", v*v)
	}

}
