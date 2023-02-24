package main

import "fmt"

func main() {
	//numbers := [5]int{5, 6, 7, 8, 9}

	//var colours [2]string = [2]string{"blue", "yellow"}
	arr := [5]int{1, 2, 3, 4, 5} // Iterated over the for loop to create an array
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//fmt.Println(colours)
	//fmt.Println(numbers)
	//fmt.Println(len(numbers))
	//fmt.Println(numbers[2])

}
