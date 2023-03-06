/*
Go program to find the sum of minimum and maximum elements.
Input:
1 2 3 4 5
Output:
min = 10
max = 14
Explanation :
1+2+3+4 = 10 for min
2+3+4+5 = 14 for max
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := make([]int, 5)
	fmt.Scanln(&a[0], &a[1], &a[2], &a[3], &a[4])
	sort.Ints(a)
	min := a[0] + a[1] + a[2] + a[3]
	max := a[1] + a[2] + a[3] + a[4]
	fmt.Println("min sum :", min)
	fmt.Println("max sum :", max)
}
