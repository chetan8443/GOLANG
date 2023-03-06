package main

import "fmt"

func main(){
	var x,y,z int=33,86,59
	a:=z-x
	b:=z-y
	if a<0{
		a=a*-1
	}
	if b<0{
		b=b*-1
	}
	print(a,b)
	if a<b{
		fmt.Println("a eats mouse first")
	}else if b<a{
		fmt.Println("b eats mouse first")
	}else{
		fmt.Println("mouse escapes")
	}
}