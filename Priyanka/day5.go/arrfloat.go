package main

import "fmt"

func main() {

	numbers := [5]float64{1.0, 2.0, 3.0, 4.0, 5.0}

	sum := 0.0
	}

	var average float64
	if len(numbers) > 0 {
		average = sum / float64(len(numbers))
	} else {
		average = 0
	}

	fmt.Printf("Average: %f\n", average)
}
