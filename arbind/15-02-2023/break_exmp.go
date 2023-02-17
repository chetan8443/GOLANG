//use of break statement in loop
package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		fmt.Printf("The value   of a is %d\n", a)
		a++
		if a > 15 {
			break
		}
	}

}
