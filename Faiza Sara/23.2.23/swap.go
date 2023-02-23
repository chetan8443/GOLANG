package main

import (
	"fmt"
	s "sara/first1"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	c, d := s.Swap(a, b)
	fmt.Println("after swaping", c, d)
}
