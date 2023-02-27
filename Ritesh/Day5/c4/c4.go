//4.Write a program that declares a slice of integers and removes all even numbers from the slice. Print out the
//resulting slice.

package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 6, 7}
	b := []int{}
	for i := range a {
		if a[i]%2 != 0 {
			b = append(b, a[i])
		}
	}
	fmt.Print(b)
}
