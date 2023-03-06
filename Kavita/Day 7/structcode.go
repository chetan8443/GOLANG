package main

import (
	"fmt"
)

type employee struct {
	name string
	age  int
	ID   int
}

//func main() {

//st := employee{"meera", 25, 100}

//fmt.Printf("%+v", st)
//}

type rectangle struct { // accessing the  elements from a structure
	l int
	b int
	//ar   int
}

func main() {
	var x rectangle
	x.l = 5
	x.b = 5

	fmt.Println((x.l) * (x.b))
}
