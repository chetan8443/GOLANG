package main

import "fmt"

func check(arr []int,element int)bool{
	var i int
	for i=0;i<len(arr);i++{
		if arr[i]==element{
			return false
		}else{
			continue
		}
	}
	return true
}

func main(){
	var arr1=[...]int{1,2,3,4,5,6,1}
	var arr2=[...]int{1,4,5,6,7,8,9,4}
	//removing duplicate in both arrays
	arr_res1:=make([]int,0)
	arr_res2:=make([]int,0)
	for _,i :=range arr1{
		if check(arr_res1,i){
			arr_res1=append(arr_res1,i)
		}else{
			continue
		}

	}
	for _,i :=range arr2{
		if check(arr_res2,i){
			arr_res2=append(arr_res2,i)
		}else{
			continue
		}

	}
	var i,j int
	fmt.Println(arr_res1,arr_res2)
	resultant_array:=make([]int,0)
	for i=0;i<len(arr_res1);i++{
		for j=0;j<len(arr_res2);j++{
			if arr_res1[i]==arr_res2[j]{
				resultant_array=append(resultant_array,arr_res1[i])
			}else{
				continue
			}
		}
	}
	fmt.Println("The intersection of two arrays is:",resultant_array)


}