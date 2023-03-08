package main

import ("fmt"
"math/rand")

func main(){
	var n,r int
	fmt.Println("Enter n")
	fmt.Scanln(&n)
	fmt.Println("Enter range")
	fmt.Scanln(&r)
	arr:=make([]int,0)
	var i int
	for i=0;i<n;i++{
		arr=append(arr,rand.Intn(r))
	}
	fmt.Println(arr)
	//find maxiumum number in a arr
	var maximum int=0
	for i=0;i<n;i++{
		if maximum<arr[i]{
			maximum=arr[i]
		}
	}
	fmt.Println(maximum)
	//count the number of maximum numbers
	var count int=0
	for i=0;i<n;i++{
		if maximum==arr[i]{
			count=count+1
		}
	}
	fmt.Println(count)
}