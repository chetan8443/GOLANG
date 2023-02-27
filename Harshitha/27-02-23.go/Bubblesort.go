package main

import "fmt"

func main(){
	var arr =[...]int{2,1,3,9,0,4,7}
	var temp int=0
	var i,j int=0,0
	for i=0;i<len(arr);i++{
		for j=0;j<len(arr);j++{
			if arr[i]<arr[j]{
				temp= arr[i]
				arr[i]=arr[j]
				arr[j]=temp
   
			}
		}
	}
	fmt.Println(arr)

	
}