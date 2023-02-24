package main

import "fmt"

func main() {
	arr := [3]float64{2.5, 3.43, 4.2}
	var average float64 = 0
	var sum float64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(float64( len(arr)))

	
	average=sum/float64( len(arr))
	fmt.Println(average)
}
