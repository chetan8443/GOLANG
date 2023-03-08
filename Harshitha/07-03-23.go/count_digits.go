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
	for n>0{
		n=n/10
		count=count+1
	}
	fmt.Println(count)

}