//Write a program that declares a slice of integers and removes all even numbers from the slice. Print out the
// resulting slice.

package main


import "fmt"

func rem(){
	num:=[]int{10,13,14,17,19,18,20,80,87,89}
	var num1[10]int

	for i:=range num{
		if(num[i]%2 != 0){

			num1[i]=num[i]
		}
	}
	fmt.Println(num1)
}