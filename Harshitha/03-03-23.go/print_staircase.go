package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Println("enter n")
	fmt.Scanln(&n)
	var i int = 0
	var k int = n
	for i = 1; i <= n; i++ {
		fmt.Println(strings.Repeat(" ", k) + strings.Repeat(" ", i+1))
		fmt.Println('\n')
	}

}
