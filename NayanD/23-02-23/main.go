package main

import (
	a "example/pack1"
	"fmt"
)

func main() {
	fmt.Println(a.Name) // global vartiable from pack1
	a.Swap()            // function from pack1
}
