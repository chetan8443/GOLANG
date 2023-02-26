package main

import (
	"fmt"
	"sort"
)

func sortAlphabetically() {
	arr := []string{"hello", "hey", "bye", "hurray", "golang", "cat", "mat", "rat"}
	slice := arr[1:7]

	fmt.Println("Before:", slice)
	sort.Strings(slice)
	fmt.Println("After:", slice)

}

func main() {
	sortAlphabetically()
}
