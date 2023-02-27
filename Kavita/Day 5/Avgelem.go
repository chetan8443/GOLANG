package main

import "fmt"

func Average(arr []float32) float32 {
	var sum float32 = 0
	var avg float32 = 0
	var n int = len(arr)
	for i := 0; i < n; i++ {

		sum = sum + arr[i]
		avg = sum / float32(n)
	}
	return avg

}

func main() {
	fmt.Println(Average([]float32{10, 10, 10, 10}))
}
