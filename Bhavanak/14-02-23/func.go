// Using Function
package main

import "fmt"

func calc(x int, y int) (int, int) {
	var out1 = x + y
	var out2 = x - y
	return out1, out2
}
func main() {
	num1 := 5
	num2 := 4
	result1, result2 := calc(num1, num2)
	fmt.Println(result1, result2)
}
