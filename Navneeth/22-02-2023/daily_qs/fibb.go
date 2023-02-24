// this program prints the fibbonacci series

package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number.")
	fmt.Scanf("%d", &num)
	a := 0
	b := 1
	temp := b
	fmt.Printf("Series: %d %d", a, b)
	for a < num {
		temp = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = temp
		fmt.Printf(" %d", b)
	}
}
