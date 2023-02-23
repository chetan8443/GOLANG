// multification table
package main

import "fmt"

func main() {
	var a int
	fmt.Println("Enter the number:")
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
