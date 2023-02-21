package main

import (
	"fmt"
	p "new/prime"  // importing external package
)

func main() {
	fmt.Print("Enter the number : ")
	n := 0
	fmt.Scan(&n)

	//calling function from imported package
	p.Prime(n)
}
