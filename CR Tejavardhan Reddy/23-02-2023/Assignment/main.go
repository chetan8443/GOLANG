package main

import (
	"fmt"
	a "tejavardhan/teja"
)

var x int = 10 //global decleration

func main() {
	fmt.Println(x)
	tejareddy()
	fmt.Println(a.SquareRoot(x))
}

func tejareddy() {
	fmt.Println(x + 10)
}
