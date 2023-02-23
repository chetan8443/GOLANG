//Nested if stmt pgm to find the biggest among three nos.

package main

import "fmt"

//Main func
func main() {
	var a,b,c ,big int
	fmt.Println("Enter 3 nos")
	fmt.Scan(&a,&b,&c)

	//nested if structure
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
	fmt.Print("Biggest number among ",a, b, c," is ",big)
}