package main

import "fmt"

func swap(px, py *int) {
	tempx := *px
	tempy := *py
	*px = tempy
	*py = tempx
}
func main() {
	x := int(1)
	y := int(2)
	fmt.Println("x was", x)
	fmt.Println("y was", y)
	swap(&x, &y)

	fmt.Println("x is now", x)
	fmt.Println("y is now", y)
}
