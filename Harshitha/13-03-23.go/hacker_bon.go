/*Question link:https://www.hackerrank.com/challenges/bon-appetit/problem?isFullScreen=true*/
package main

import "fmt"

func main(){
	var n,k int
	fmt.Println("enter n")
	fmt.Scanln(&n)
	fmt.Println("Enter k")
	fmt.Scanln(&k)//which anna doesnot eat
	bill:=make([]int,n)
	var i int=0
	//Enter elements into the array
	fmt.Println("Enter Elements")
	for i=0;i<n;i++{
		fmt.Scanln(&bill[i])
	}
	fmt.Println(bill)
    var paid int
	fmt.Println("enter the amount paid by anna")
	fmt.Scanln(&paid)
	var sum int
	for i=0;i<n;i++{
		if i!=k{
			sum=sum+bill[i]
		}
	}
	if sum/2==paid{
		fmt.Println("Bon Appetite")
	}else{
		fmt.Println(paid-(sum/2))
	}


}
