package main

import "fmt"

func main() {
	mymp := map[int]int{1: 6, 2: 7, 3: 8, 4: 9}

	for key, value := range mymp {

		fmt.Println(key, value)
	}

}
