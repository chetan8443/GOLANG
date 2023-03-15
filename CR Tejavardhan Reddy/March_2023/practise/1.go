package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Println("Enter the number:")
	fmt.Scanln(&num)
	res, err := Sqrt(num)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(res)
}

func Sqrt(val float64) (float64, error) {
	if val < 0 {
		return 0, errors.New("Math:Negative number passed")
	}
	return math.Sqrt(val), nil
}
