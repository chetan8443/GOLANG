package main

import "fmt"

func main() {
	var i int

	var arr1 [5]int
	var arr2 [5]int
	var Marr [5]int

	fmt.Print("Enter First array elements")
	for i = 0; i < 5; i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Enter Second array elements ")
	for i = 0; i < 5; i++ {
		fmt.Scan(&arr2[i])
	}

	fmt.Print("Multiplication")
	for i = 0; i < len(Marr); i++ {
		Marr[i] = arr1[i] * arr2[i]
		fmt.Print(Marr[i], "  ")
	}
	fmt.Println()
}
