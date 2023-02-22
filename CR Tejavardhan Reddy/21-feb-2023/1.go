/*
Pattern Printing
Input:5
Output:
1
1 2
1 2 3
1 2 3 4
1 2 3 4 5
1 2 3 4
1 2 3
1 2
1
*/
package main

import (
	"fmt"
)

func main() {
	//fmt.Println(S)
	fmt.Print("Enter the number:")
	var n int
	fmt.Scanln(&n) //reading the range
	for i := 1; i < n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
	for k := n; k > 0; k-- {
		for l := 1; l <= k; l++ {
			fmt.Print(l)
			fmt.Print(" ")
		}
		fmt.Print("\n")
	}
}
