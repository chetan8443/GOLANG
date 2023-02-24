package main

import "fmt"

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Before Reversing")
	fmt.Println(list) // printing before reversing a slice

	var new_list []int // declaring empty  slice
	for i := len(list) - 1; i >= 0; i-- {
		new_list = append(new_list, list[i])
	}
	fmt.Println("After Reversing")
	fmt.Println(new_list) // printing after reversing a slice
}
