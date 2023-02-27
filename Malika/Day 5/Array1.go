//Write a program that declares an array of integers and prints out the sum of the elements.//

package main

import "fmt"

func array(arr []int) {

	var sum int = 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	fmt.Println(sum)
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	array(arr)
}
