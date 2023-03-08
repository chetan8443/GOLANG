package main

import "fmt"

func main() {
	var arr = []string{"aaa", "bbb", "ccc", "bbb", "aaa", "aaa"}
	var list = make([]int, 0)
	var dict = make(map[string]int)
	for _, num := range arr {
		dict[num] += 1
	}
	fmt.Print(dict)
	for _, v := range dict {
		list = append(list, v)
	}

}
