package main

import ("fmt")

func main() {
	var a int = 100
	var b int = 200
	var ret int
	ret = min(a, b)
	fmt.Printf("Minimum value is : %d\n", ret)
}
func min(num1, num2 int) int {
	var result int
	if num1 < num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}