package main

import "fmt"

func main() {
	var i int
	var arr1 = [5]int{1, 2, 3, 4, 5}
	var arr2 = [5]int{6, 7, 8, 9, 10}
	fmt.Println("Add\t Sub\t Mul\t Div\t mod\t")
	for i = 0; i < 5; i++ {
		fmt.Print("\n", arr1[i]+arr2[i], "\t")
		fmt.Print(arr1[i]-arr2[i], "\t")
		fmt.Print(arr1[i]*arr2[i], "\t")
		fmt.Print(arr1[i]/arr2[i], "\t")
		fmt.Print(arr1[i]%arr2[i], "\t")
	}
}
