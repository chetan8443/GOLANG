// write a program to print pyramid star pattern.

package main

import "fmt"

func main() {
	var rows int = 5
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}
