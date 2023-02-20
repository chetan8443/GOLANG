package main

import "fmt"

func main() {
	var i, j, rows int
	fmt.Println("Enter number of rows=")
	fmt.Scanln(&rows)
	for i = 1; i <= rows; i++ {
		for j = 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
