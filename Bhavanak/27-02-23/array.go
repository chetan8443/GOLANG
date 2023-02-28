package main

import "fmt"

func main() {
	var i int
	var arr [5]int
	for i = 0; i < 5; i++ {
		fmt.Println("enter a number")
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	var res [5]int
	var k int = 0
	for i = 4; i >= 0; i-- {
		res[i] = arr[k]
		k = k + 1
	}
	fmt.Println(res)
}
