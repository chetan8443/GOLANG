package main

import "fmt"

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	sum := 0
	for val := range arr {
		sum += val
	}
	fmt.Print("sum of the elements : ", sum)
}
