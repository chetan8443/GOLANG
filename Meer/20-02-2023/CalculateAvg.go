package main

import "fmt"

func main() {
	var sum, avg float64
	var temp int
	fmt.Print("Enter number of elements: ")
	fmt.Scanln(&temp)
	var num = make([]float64, temp)
	for i := 0; i < temp; i++ {
		fmt.Print("Enter the number : ")
		fmt.Scanln(&num[i])
		sum += num[i]
	}

	avg = sum / float64(temp)
	fmt.Printf("The Average of entered %d number(s) is %f", temp, avg)
}
