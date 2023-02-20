// Star Pyramid Pattern

package main

import "fmt"

func main() {

	var i, j, rows int

	fmt.Print("Enter number of rows:= ")
	fmt.Scanln(&rows)

	for i = 1; i <= rows; i++ {
		for j = 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k != (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
