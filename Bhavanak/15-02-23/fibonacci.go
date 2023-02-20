// Check the given number is fibonacci series using function

package main

import "fmt"

func fibonacci(n int) {
	var f3, f2, f1 int = 0, 0, 1
	for i := 1; i <= n; i++ {
		fmt.Println(f1)
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
}

func main() {
	var n int
	fmt.Println("Enter the range of fibonacci numbers:")
	fmt.Scanln(&n)
	fibonacci(n)
}
