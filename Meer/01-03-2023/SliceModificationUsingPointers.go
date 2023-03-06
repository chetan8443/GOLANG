package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	modifySlice(&slice1)
	fmt.Println(slice1)
}

func modifySlice(slice1 *[]int) {
	*slice1 = append((*slice1), 6, 7, 8, 9)
}
