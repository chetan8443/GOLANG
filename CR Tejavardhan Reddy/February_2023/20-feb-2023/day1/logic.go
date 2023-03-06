package teja

import (
	"math"
)

func PerfectSquare(n float64) bool {
	if int(n) == 1 || int(n) == 0 {
		return true
	} else {
		a := math.Sqrt(n)
		if a*a == n {
			return true
		} else {
			return false
		}

	}
	//fmt.Print("Is a perfect square")
}
