package main

import "fmt"

func main() {
	rows := 5
	for i := rows; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			fmt.Printf(" ")
		}
		for k := i; k <= (2*i - 1); k++ {
			fmt.Printf("*")
		}
		for k := 0; k < i-1; k++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}

}
