// write a code to find average of numbers present in a slice

package main

import "fmt"

func main() {

	list := []int{11, 33, 22, 44, 33, 22, 44}
	fmt.Printf("given list is %v \n", list)
	fmt.Printf("the average of given list is %v \n", avg(list))

}

func avg(list []int) float64 {
	sum := 0
	count := 0
	for i := 0; i < len(list); i++ {
		sum += list[i]
		count++
	}
	avg := float64(sum) / float64(count)
	return avg
}
