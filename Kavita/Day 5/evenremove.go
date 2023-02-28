package main

import "fmt"

func main() {

	slc := []int{10, 11, 30, 5, 7}
	for i := 0; i < len(slc); i++ {
		if slc[i]%2 == 0 {
			continue
		} else {
			fmt.Println(slc[i])
		}
	}
}
