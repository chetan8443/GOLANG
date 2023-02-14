package main

import "fmt"

func main() {
	fmt.Print("please eneter no")
	var count int
	fmt.Scan(&count)

	var a int = 0
	fmt.Println(a)
	var b int = 1
	fmt.Println(b)
	for i := 0; i < count; i++ {

		var c int = a + b
		if(c<count){
			fmt.Println(c)
			a=b
			b=c
		}
	
	}
}
