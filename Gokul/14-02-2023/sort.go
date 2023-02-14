//Program to sort an array using bubble sort
package main

import "fmt"

func main() {
	var size int
	fmt.Println("Enter size of array")
	fmt.Scan(&size)
	var arr =make([]int,size) 		//creates array of int type with the required size only
	fmt.Println("Enter elements of array")
	for i:=0;i<=size-1;i++{
		fmt.Scan(&arr[i])
	}
	for i:=0;i<size-1;i++{
		for j:=0;j<size-i-1;j++ {
			if arr[j] >arr[j+1]{
				arr[j],arr[j+1] =arr[j+1],arr[j]  //swapping
			}
		}
	}
	fmt.Println("Sorted array is ",arr)
}