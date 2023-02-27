package main

import "fmt"

func main() {
	var size int
	fmt.Print("Enter the size:")
	fmt.Scanln(&size)
	var arr = make([]int, size)
	fmt.Println("Enter the numbers:")

	for i := range arr {
		// fmt.Println(i)
		// fmt.Printf("enter ith element")
		fmt.Scanf("%d\n", &arr[i])
	}
	fmt.Println("The array is :", arr)
	var revArray []int
	for i := size - 1; i >= 0; i-- {
		revArray = append(revArray, arr[i])
	}
	fmt.Println("Reversed array:", revArray)

}
