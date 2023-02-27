package main

import (
	a "dev/test"
)

func main() {
	arr := [5]int{2, 3, 4, 5, 6}
	a.Sum(arr)
	var arr1 = [5]float64{1.2, 3.4, 4.5, 6.7, 8.0}
	a.Avg(arr1)
	var arr2 = []string{"e", "d", "c", "b", "a"}
	a.Sort(arr2)

	var arr3 = []int{2, 3, 4, 5, 6, 7, 8}
	a.Even_remove(arr3)
}
