package main

import "fmt"

func main(){
	var n,i int
	fmt.Println("enter the size of the array")
	fmt.Scanln(&n)
	var hars [5] int
	for i=0;i<5;i++{
		fmt.Println("enter a number")
		fmt.Scanln(&hars[i])
	}
	for i=0;i<5;i++{
		fmt.Println(hars[i])
	}
}