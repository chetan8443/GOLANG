package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter the size:")
	var size int
	fmt.Scanln(&size)
	var arr = make([]int, size)
	for i := 0; i < size; i++ {
		fmt.Scanln(&arr[i])
	}
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	initial := arr[0]
	//fmt.Print(initial)
	badRecords := 0
	goodRecords := 0
	for num := range dict {
		//fmt.Println(num)
		if num > initial {
			goodRecords += 1
		} else if num < initial {
			badRecords += 1
		}
	}
	fmt.Print("Number of bad recors:", badRecords)
	fmt.Print("Number of good records:", goodRecords)
}
