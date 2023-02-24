package main

import (
	a "new/module"
)

func main() {

	var arr = [5]int{1, 2, 3, 4, 5}
	a.Sum(arr)

	var arrF = [5]float64{1.2, 2, 4, 5.5, 5.9}
	a.Avg(arrF)

	var s1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a.RemEven(s1)

	var str = []string{"aniket", "khupte", "chetan"}
	a.Sort(str)
}
