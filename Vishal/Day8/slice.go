
package main

import "fmt"

func main() {
	fmt.Println("")
	// declaring slice with some integer numbers
	var s1= []int {1,2,12,43,21,34,40,25}
    ptr:=&s1   // address of s1 slice assigned to pointer
	var even=[]int{}
	var odd =[]int{}
    var ptr1 =&even
	var ptr2 =&odd
		for _, v := range *ptr {
		if v%2==0 {//checking odd even numbers
			*ptr1=append(*ptr1, v)   // appending slice using pointer
		}else {
			*ptr2=append(*ptr2, v)
		}
	}
// printing both slices using their identifiers
	fmt.Println("Even Numbers ",even)
	fmt.Println("odd Numbers",odd)
}
