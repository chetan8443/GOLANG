//Write a program that declares an array of integers and prints out the sum of the elements.

package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter the size of the array:")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Println("Enter the elements:")
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	//fmt.Println(arr)
	sum := 0
	for _, element := range arr {
		sum += element
		//fmt.Println(i)
	}
	fmt.Print("The sum of the elements:", sum)

}
