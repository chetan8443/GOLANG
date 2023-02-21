package main

import "fmt"

func main() {
	slice1 := make([]int, 5)
	fmt.Print("Give numbers to make a slice :")
	var num int

	//for loop for take numbers in a slice
	for i := 0; i < len(slice1); i++ {
		fmt.Scan(&num)
		slice1[i] = num
	}

	//print slice
	for i, s := range slice1 {
		fmt.Println(i, s)
	}

	fmt.Print("Type number which want to search : ")
	var num2 int
	fmt.Scan(&num2)

	//for loop for search index of number
	for i := 0; i < len(slice1); i++ {
		if slice1[i] == num2 {
			fmt.Printf("given %d number found in %dth index", num2, i)
		} else {
			fmt.Print("number is not found")
		}
	}

}
