package main

import (
	"fmt"
	
)

func main() {
	var n int
	fmt.Print("Ente a num : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		a:=0
		b:=1
		fmt.Print(b)
		for j := 0; j < i; j++ {
			c:=a+b
			fmt.Print(c)
			a=b
			b=c

		}
		fmt.Println("")
	}
}
