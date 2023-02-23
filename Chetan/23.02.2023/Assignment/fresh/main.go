package main

import ("fmt"
		a "hi/fc")

func main() {
	// R:=a.Sum(4,6)
	m:=a.Multi(500,10)
	if (m<100){
	
		fmt.Printf("sum is %d less than 100",m)
	}else{
		fmt.Printf("sum is %d greater than 100",m)
	}
}