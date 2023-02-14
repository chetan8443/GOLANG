package main

import "fmt"

func main() {
	var a int = 1
	var x int = 6
	var i int = 1
	for i = 1; i <= x; i++ {
		a = a * i
	}
	fmt.Println(a)

}
