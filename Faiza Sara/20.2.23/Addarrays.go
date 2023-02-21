package main

import "fmt"

func main() {
	var s, i int

	fmt.Print("Enter the Array Size")
	fmt.Scan(&s)

	arr1 := make([]int, s)
	arr2 := make([]int, s)
	addarr := make([]int, s)

	fmt.Print("Enter the first array elements")
	for i = 0; i < s; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Enter the second array elements")
	for i = 0; i < s; i++ {
		fmt.Scan(&arr2[i])
	}

	fmt.Print("The Sum of two Arrays")
	for i = 0; i < s; i++ {
		addarr[i] = arr1[i] + arr2[i]
		fmt.Print(addarr[i])
	}
	fmt.Println()
}
