package main

import "fmt"

func main(){
	var n int
	fmt.Println("enter n")
	fmt.Scanln(&n)
	divisors:=make([]int,0)
	var i int
	for i=1;i<n;i++{
		if n%i==0{
			divisors=append(divisors,i)
		}
	}
	fmt.Println(divisors)
	var j,sum int=0,0
	for j=0;j<len(divisors);j++{
		sum=sum+divisors[j]
	}
	fmt.Println(sum)
	if sum==n{
		fmt.Println("The given number is a perfect number")
	}else{
		fmt.Println("The given number is not a perfect number")
	}
}