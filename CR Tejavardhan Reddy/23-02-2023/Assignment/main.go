package main

import (
	"fmt"
	a "tejavardhan/teja"
)

var x int = 10 //global decleration

func main() {
	fmt.Println(x)
	tejareddy()
	fmt.Println(a.SquareRoot(x)) //Accessing the another file function
}

func tejareddy() {
	fmt.Println(x + 10)
	c := 0
	for i := 0; i < x; i++ {
		if x/i == 0 {
			c++
		} else {
			continue
		}
	}
	fmt.Println("Enter a number:")
	var num int
	fmt.Scanf("%d", num)
	switch num { //switch case in functions
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	default:
		fmt.Println("Error")
	}
}
