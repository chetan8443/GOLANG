package main

import "fmt"

var flag int = 0

func main() {
	var index int
	fmt.Print("Enter the size of the array:")
	var n int
	fmt.Scanln(&n)
	var arr = make([]int, n)
	fmt.Println("Enter the elements into the array:")
	for i := 0; i < n; i++ {

		fmt.Scanln(&arr[i])
	}
	fmt.Print("Enter the number you want to search:")
	var target int
	fmt.Scanln(&target)
	for i, ele := range arr {
		if ele == target {
			flag = 1
			index = i
		}
	}
	if flag == 0 {
		fmt.Print("The element is not found")
	} else {
		fmt.Print("The element is at ", index, " index")
	}

}
