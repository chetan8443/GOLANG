package main

import "fmt"

func arraysum(arr [] int) int {
	res:=0
	for i:=1;i<len(arr);i=i+2{
		res=res+arr[i]
	}
	return res
}

func main(){
	fmt.Println(arraysum([]int{1,2,3,4,5}))

}