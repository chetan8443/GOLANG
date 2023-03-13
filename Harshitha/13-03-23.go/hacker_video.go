//Question Link:https://www.hackerrank.com/challenges/the-hurdle-race/problem?isFullScreen=true
package main

import "fmt"

func main(){
	var n int
	fmt.Println("Enter the value of n")
	fmt.Scanln(&n)
	heights:=make([]int,n)
	fmt.Println("Enter elements into the array")
	var i int
	for i=0;i<n;i++{
		fmt.Scanln(&heights[i])
	}
	fmt.Println(heights)
	/*5 4
1 6 3 5 2*/
    //find the maximum of the array
	maximum:=0
	for i=0;i<n;i++{
		if heights[i]>maximum{
			maximum=heights[i]
		}
	}
	var maxi int
	fmt.Println("enter the maximum height a hurdle can take up")
	fmt.Scanln(&maxi)
	if maximum>maxi{
		fmt.Println("It mmust drink two units more to jump through all the hurdles",maximum-maxi)
	}else{
		fmt.Println("It can jump through all the hurdles without any disturbance")
	}
}