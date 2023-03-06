package main

import "fmt"
func check1(arr []int,element,j,n,k int)(bool,int){
	var i int
	for i=j;i<n;i++{
		if arr[i]==element{
			return true,i
		}else{
			continue
		}
	}
	return false,0
}

func main(){
	var n int
	var k int
	fmt.Println("Enter numbers n,k")
	fmt.Scanln(&n,&k)
	arr:=make([]int,n)
	var i int
	fmt.Println("enter the elements into the array")
	for i=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	fmt.Println(arr)
	var a bool
	var b int
	for i=0;i<n;i++{
		element:=k-arr[i]
		a,b=check1(arr,element,i,n,k)
		if a{
			fmt.Println(arr[i],arr[b])
			fmt.Println("......")
			}
	}

}