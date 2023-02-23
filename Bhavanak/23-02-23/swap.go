package main

import (
	t "bhav/Assignment"
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("enter a value:")
	fmt.Scanln(&a)
	fmt.Println("enter b value:")
	fmt.Scanln(&b)
	c, d := t.Swap(a, b)
	fmt.Println("after swapping", c, d)

}
