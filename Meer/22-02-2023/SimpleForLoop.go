package main

import (
	"fmt"
)

func main() {
	var fruitlist = []string{"Apple", "mango", "grapes"}
	fruitlist = append(fruitlist, "strawberry", "peach", "banana")

	for i, fruite := range fruitlist {
		fmt.Println(i+1, " ", fruite)
	}

}
