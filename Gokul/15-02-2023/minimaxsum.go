//Program to find the min and max sum of pair of elements in a slice
package main

import (
	"fmt"
	"sort"
)

func main() {
	var size int
	fmt.Println("Enter size of array")
	fmt.Scan(&size)
	var ar =make([]int,size)
	fmt.Println("Enter elements of array")
	for i:=0;i<size;i++ {
		fmt.Scan(&ar[i])
	}
	sort.Ints(ar)
	minsum := ar[0] + ar[1]
	maxsum:= ar[size-1] + ar[size-2]
	fmt.Print("Min sum is of pair ",ar[0],", ",ar[1]," and is ",minsum,"\n")
	fmt.Print("Max sum is of pair ",ar[size-1],", ",ar[size-2]," and is ",maxsum,"\n")
}