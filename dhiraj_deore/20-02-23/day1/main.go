package main

import (
	"fmt"
	p "new/prime"
)

func main() {
	fmt.Print("Enter the number : ")
	n := 0
	fmt.Scan(&n)

	//Prime(n)
	p.Prime(n)
}
