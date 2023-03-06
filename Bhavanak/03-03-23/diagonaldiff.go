/*
Go program to find the difference between sums of two .

input : mat[][] = 11 2 4

	 4 5 6
	10 8 -12

Output : 15
Sum of primary diagonal = 11 + 5 + (-12) = 4.
Sum of secondary diagonal = 4 + 5 + 10 = 19.
Difference = |19 - 4| = 15.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int = 3
	var a = [3][3]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}

	d1 := 0
	d2 := 0

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				d1 += a[i][j]
			}
			if i+j == n-1 {
				d2 += a[i][j]
			}
		}
	}

	diff := math.Abs(float64(d1 - d2))
	fmt.Println(diff)

}
