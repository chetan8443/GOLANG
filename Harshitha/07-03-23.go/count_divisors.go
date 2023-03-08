package main

import ("fmt"
"math/rand")

func main(){
	var Range int
    fmt.Println("Enter the range :")
    fmt.Scanln(&Range)
    n:= rand.Intn(Range)
	fmt.Println(n)
	count:=0
	a:=n
	res:=0
	for n>0{
		res=n%10
		if a%res==0{
			count=count+1
		}
		n=n/10
	}
	fmt.Println(count)
}