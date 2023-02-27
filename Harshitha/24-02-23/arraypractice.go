package main

import ("fmt"
"math/rand")

/*func main(){
	var arr=[...]int{1,2,3,4,5,6}
	fmt.Println(len(arr))
	fmt.Println(arr)
	arr1:=arr
	arr1[0]=10
	fmt.Println(arr1)
	if arr1==arr{
		fmt.Println("Both the arrays are equal")
	}else{
		fmt.Println("Not equal")
	}
	var sum int=0
	var i int=0
	for i=0;i<len(arr);i++{
		sum=sum+arr[i]
	}
	fmt.Println("The sum of elements of array arr is:",sum)
	var maximum int=0
	for i=0;i<len(arr1);i++{
		if maximum<arr1[i]{
			maximum=arr1[i]
		}
	}
	fmt.Println("The maximum number in the array is",maximum)
	var even = make([]int,len(arr))
	var odd =make([]int,len(arr))
	var e int=0
	var o int=0
	for i=0;i<len(arr);i++{
		if arr[i]%2==0{
			even[e]=arr[i]
			e=e+1
		}else{
			odd[o]=arr[i]
			o=o+1
		}
	}
	fmt.Println(even)
	fmt.Println(odd)
	var c =make([]int,len(arr)+len(arr1))
	var k int=0
	for i=0;i<len(arr);i++{
		c[k]=arr[i]
		k=k+1
	}
	for i=0;i<len(arr1);i++{
		c[k]=arr1[i]
		k=k+1
	}
	fmt.Println("The merged array is",c)
	//add two integer arrays
	j:=0
	for i=0;i<len(arr);i++{
		c[j]=arr[i]+arr1[i]
		j=j+1
	}
	fmt.Println("The sum of two arrays is:",c)
	//usage of map function in golang
	var a string="apple banana"
	var b string="banana"
	h=a.Split(a," ")
	t=b.Split(b," ")
	func search(element) bool{
		for i=0;i<len(counter);i++{
			if counter[i]==element{
				return true
			}else{
				return false
			}
		}

	}
	var counter = map[string] int {}
	for i=0;i<len(h);i++{
		if search(h[i]){
			counter[i]=counter[i]+1
		}
	}
	for i=0;i<len(t);i++{
		if search(t[i]){
			counter[i]=counter[i]+1
		}
	}*/
func main(){
	randominteger:=rand.Intn(100-0)+0
	fmt.Println(randominteger)
} 