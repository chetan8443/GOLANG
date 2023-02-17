package main

import "fmt"

func main() {
	slice1 := []string{"Infobell", "IT"}
	slice2 := []string{"Infobell", "IT"}

	fmt.Printf("slices are %v, %v and these are equal: %t\n", slice1, slice2, compareSlices(slice1, slice2))

	slice2 = append(slice2, "Solutions Pvt., Ltd.")
	fmt.Printf("slices are %v, %v and these are equal: %t\n", slice1, slice2, compareSlices(slice1, slice2))
}

func compareSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, strs := range slice1 {
		if strs != slice2[i] {
			return false
		}
	}
	return true
}
