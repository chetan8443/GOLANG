//Write a program to reverse a slice

package main

import "fmt"

func main() {
	slice := []int{12, 32, 45, 21, 34, 43}
	rev := make([]int, 6)
	fmt.Println("Slice before reverse numbers :", slice)

	for i := 0; i < len(slice); i++ {
		rev[i] = slice[len(slice)-i-1]
	}

	fmt.Print(rev)

}
