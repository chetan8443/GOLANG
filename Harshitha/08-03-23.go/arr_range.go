package main

import "fmt"
func counter(arr []int,element int)int{
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
	var n,i int
	fmt.Println("Enter the no.of elements in the array")
	fmt.Scanln(&n)
	arr:=make([]int,n)
	fmt.Println("THe elements in the array are")
	for i=0;i<n;i++{
		fmt.Scanln(&arr[i])
	}
	result_array:=make([]int,0)
	fmt.Println(arr)
	//To find the maximum of an array
	var maximum int=0
	for i=0;i<n;i++{
		if maximum<arr[i]{
			maximum=arr[i]
		}
	}
	//To count values

	for i=1;i<maximum+1;i++{
		a:=counter(arr,i)
		result_array=append(result_array,a)
	}
	fmt.Println(result_array)
	
}