/*
User enters an array
Then a number is entered by the user
If the number is present in the array then the index is returned
else -1 is returned
Input1:[2,35,7,6,9,0]
7
Output1: 2
***********************
Input2:[2,35,7,6,9,0]
8
Output2:-1
*/
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
		fmt.Print("-1")
	} else {
		fmt.Print("The element is at ", index, " index")
	}

}
