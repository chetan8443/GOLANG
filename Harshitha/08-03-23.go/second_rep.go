package main

import "fmt"
//keys can not be duplicated in map

func counter1(arr []string,element string)int{
	var count int=0
	var i int
	for i=0;i<len(arr);i++{
		if element==arr[i]{
			count=count+1
		}
	}
	return count
}

func main(){
	var n int
	fmt.Println("Enter the number of elements into array")
	fmt.Scanln(&n)
	arr:=make([]string,n)
	var i int
	fmt.Println("Enter the elements into the array")
	for i=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	var result = map[int]string{}
	var a int
	for i=0;i<n;i++{
		a=counter1(arr,arr[i])
		result[a]=arr[i]
	}
	
	fmt.Println(result)
}