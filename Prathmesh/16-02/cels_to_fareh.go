package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter the temperature in Celsius:")
	fmt.Scanf("%d", &n)
	f := (float32(n) * 1.8) + 32
	fmt.Println("Temperature in Fahrenheit is:", f)
}
