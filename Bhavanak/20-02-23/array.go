package main

import "fmt"

func main() {
	var i int
	var arr1 [5]int
	var arr2 [5]int
	var addarr [5]int
	fmt.Print("Enter the items in array :=")
	for i = 0; i < 5; i++ {
		fmt.Scan(&arr1[i])
	}
	fmt.Print("Enter the items in array :=")
	for i = 0; i < 5; i++ {
		fmt.Scan(&arr2[i])
	}
	fmt.Print("Sum of two arrays is =")
	for x, _ := range arr1 {
		addarr[x] = arr1[x] + arr2[x]
		fmt.Print(addarr[x], " ")
	}
	fmt.Println()
}
