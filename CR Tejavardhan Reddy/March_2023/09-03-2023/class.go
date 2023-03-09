package main

import "fmt"

func main() {
	fmt.Println("Enter minimum number of students:")
	var k int
	fmt.Scanln(&k)
	fmt.Println("Enter total number of students:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Println("Enter the time delay/time early by minutes of each student")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	var ontime = 0
	for _, num := range arr {
		if num <= 0 {
			ontime++
		}
	}
	if ontime >= k {
		fmt.Println("The class will continue")
	} else {
		fmt.Println("The class will be suspended")
	}
}
