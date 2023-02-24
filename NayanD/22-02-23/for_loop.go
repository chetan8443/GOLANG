// write a program based on for loops. printing inverted pyramid star pattern.

package main

import "fmt"

func main() {
	rows := 5                 // number of rows
	for i := rows; i >= 1; i-- {
		for j := 1; j <= rows-i; j++ {
			fmt.Printf(" ")               // printing spaces
		}
		for k := i; k <= (2*i - 1); k++ {
			fmt.Printf("*")
		}
		for k := 0; k < i-1; k++ {
			fmt.Printf("*")   //printing star 
		}
		fmt.Println()
	}

}
