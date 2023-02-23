// Write a program to draw a star pattern using nested for loop :

package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		for k := 6 - i; k >= 1; k-- {
			fmt.Print("   ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print(" * ")
		}

		for k := 1; k <= i; k++ {
			fmt.Print(" * ")
		}
		fmt.Println()
	}
}
