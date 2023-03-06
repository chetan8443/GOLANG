// Write a program to Print Floyd's Triangle for 5 rows.

package main

import "fmt"

func main() {
	c := 0
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			c++
			fmt.Print(" ", c)
		}
		fmt.Println("")
	}
}
