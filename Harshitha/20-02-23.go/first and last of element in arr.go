package main

import (
	"fmt"
)

func main(){
	var hars [7] int={1,2,2,2,3,4,5}
	var n int
	fmt.Println("enter the number to be searched")
	fmt.Scanln(&n)
	var first int =-1
	var last int =-1
	for i=0;i<7;i++{
		if hars[i]!=n{
			continue
		}
		if first==-1:{
			first=i
		}
		last=i

	}
}
