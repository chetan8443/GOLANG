package main

import "fmt"

func main() {
	Range := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	Slice_1 := Range[2:4]
	fmt.Println(Slice_1)
	fmt.Println(len(Slice_1))
	fmt.Println(cap(Slice_1))

}
