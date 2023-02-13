package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	i := 1
	for {
		if i > 10 {
			break
		}
		fmt.Println(a, "*", i, "=", a*i)
		i++
	}
}
