package main

import "fmt"

func main() {
	fmt.Println("Enter 3 nos")
	var a,b,c ,big int
	fmt.Scan(&a,&b,&c)
	if a > b{
		if(a>c){
			big = a
		}else{
			big = c
		}
	}else{
		if b>c {
			big = b
		}else{
			big = c
		}
	}
	fmt.Printf("Biggest number among the three is %d",big)
}