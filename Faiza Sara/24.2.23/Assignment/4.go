// Write a program that declares a slice of integers and removes all even numbers from the slice.
// Print out the resulting slice.
package main

import "fmt"

func main() {
	var s int
	fmt.Print("Enter the size of the array:")
	fmt.Scanln(&s)
	var arr = make([]int, s)
	fmt.Println("Enter the elements:")
	for i := 0; i < s; i++ {
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	for i, element := range arr {
		fmt.Println(i, element)
		if element%2 == 0 {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	fmt.Print(arr)
}
