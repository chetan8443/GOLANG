package main

import "fmt"

func main() {
	var r int
	var k int = 0
	fmt.Print("Enter number of rows")
	fmt.Scan(&r)
	for i := 1; i <= r; i++ {
		k = 0
		for s := 1; s <= r-i; s++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}
}
